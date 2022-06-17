#!/bin/zsh

# This file is used to build kubernetes-resource-information-collector tool.

function get_root_path {
    work_directory=$(cd $(dirname $0) && pwd)
    root_directory="${work_directory%/*}"
    echo "${root_directory}"
}

function error {
    message=$1
    if [[ "${message}" == "" ]]; then
        message="Unknown error"
    fi
    echo -e "[ERROR] $1"

    code=$2
    if [[ "${code}" == "" ]]; then
        code=1
    fi

    exit $code
}

function check {
    command helm > /dev/null 2>&1
    if [[ $? -ne 0 ]]; then
        error "Command not found: helm" 2
    fi

    command go > /dev/null 2>&1
    if [[ $? -ne 2 ]]; then
        error "Command not found: go" 2
    fi
}

# This function is used to inject version information to each component.
function replace_version {
    release_directory=$1

    # Replace version in collector-prerequisite chart
    sed -i '.bak' "s/version: version/version: ${COLLECTOR_CHARTS_PREREQUISITE_VERSION}/g" "${release_directory}"/collector-prerequisite/Chart.yaml
    rm -rf "${release_directory}"/collector-prerequisite/Chart.yaml.bak

    # Replace version in README.md
    sed -i '.bak' "s/# kubernetes-resource-infomation-collector version/# kubernetes-resource-infomation-collector ${VERSION}/g" "${release_directory}"/README.md
    rm -rf "${release_directory}"/README.md.bak

    # Replace version in manage.sh
    sed -i '.bak' "s/    echo VERSION/    echo ${VERSION}/g" "${release_directory}"/manage.sh
    rm -rf "${release_directory}"/manage.sh.bak

    sed -i '.bak' "s/COLLECTOR_PREREQUISITE_CHART_VERSION=version/COLLECTOR_PREREQUISITE_CHART_VERSION=${COLLECTOR_CHARTS_PREREQUISITE_VERSION}/g" "${release_directory}"/manage.sh
    rm -rf "${release_directory}"/manage.sh.bak

    # Replace version in collector
    sed -i '.bak' "s/__VERSION__/${COLLECTOR_COLLECTOR_VERSION}/g" "${release_directory}"/collector/module/service.go
    rm -rf "${release_directory}"/collector/module/service.go.bak

    # Replace version in lcweb
    sed -i '.bak' "s/__VERSION__/${WEBVERSION}/g" "${release_directory}"/lcweb/package.json
    rm -rf "${release_directory}"/lcweb/package.json.bak
}

# This function is used to copy source code(helm charts, README and go etc.) to integrate directory
function copy_source_code {
    root_path=$1
    release_directory=$2

    # Copy collector-prerequisite chart and move values-override.yaml file in release directory directly.
    cp -r "${root_path}"/charts/collector-prerequisite "${release_directory}"
    mv "${release_directory}"/collector-prerequisite/values-override.yaml "${release_directory}"/values-override-collector-prerequisite.yaml

    # Copy README.
    cp "${root_path}/README.md" "${release_directory}"

    # Copy management tool
    cp "${root_path}/manage/manage.sh" "${release_directory}"
    chmod 740 "${release_directory}/manage.sh"

    # Copy collector
    cp -r "${root_path}/collector" "${release_directory}"

    # Copy lcweb
    cp -r "${root_path}/lcweb" "${release_directory}"
}

# This function is used to delete source code after building
function delete_source_code {
    release_directory=$1

    # Delete collector-prerequisite chart without values-override.yaml
    rm -rf "${release_directory}/collector-prerequisite"

    # Delete server code
    rm -rf "${release_directory}/collector"
}

# This function used to build each components
function build_components {
    release_directory=$1

    # build collector-prerequisite
    # Notice that cannot delete this `cd` command, or it will be generate helm charts in current directory
    cd "${release_directory}"
    helm package "${release_directory}"/collector-prerequisite
    cd -

    # build log-collector server code
    cd "$release_directory"/collector/cmd
    go build -o log-collector main.go
    mv log-collector "${release_directory}/"
    cd -
}

# Main function for this script.
function main {
    # check dependency tool
    check

    root_path=$(get_root_path)

    # load version information
    version_file_path="${root_path}/kric.version"
    if [ ! -f "${version_file_path}" ]; then
        error "Can not find version file"
    fi
    source "${version_file_path}"

    # generate integrate directory
    integrate_directory="${root_path}/integrate"
    mkdir -p "${integrate_directory}"

    release_directory="${integrate_directory}/kubernetes-resource-information-collector-${VERSION}"
    rm -rf release_directory > /dev/null 2>&1
    mkdir "${release_directory}"

    # copy source code to integrate directory
    copy_source_code "${root_path}" "${release_directory}"

    # replace version for each component
    replace_version "${release_directory}"

    # build components
    build_components "${release_directory}"

    # delete source code which in release directory
    delete_source_code "${release_directory}"

    # build package
    cd "${integrate_directory}"
    tar zcvf kubernetes-resource-information-collector-"${VERSION}".tar.gz ./kubernetes-resource-information-collector-"${VERSION}"
    rm -rf kubernetes-resource-information-collector-"${VERSION}"
    cd -
}

main
