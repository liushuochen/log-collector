#!/bin/zsh

# kubernetes-resource-information-collector management tool

# ========== global variables ==========
# The log level.
# Support 1, 2, 3, 4.
# Type: integer.
# Default: 3.
# level 1: Print debug, info, warning and error log.
# level 2: Print info, warning and error log.
# level 3: Print warning and error log.
# level 4: Only print error log.
# It can be reset by reset_log_level function. Please not modify it directly.
LOG_LEVEL=3

# A flag indicate that debug mode whether opening.
# Support: true, false
# Default: false
# true: debug mode opening
# It can be reset by open_debug_mode function. Please not modify it directly.
DEBUG_MODE="false"

# A kubernetes namespace name which deploy resources under it. Default is default.
# It can be reset by reset_namespace function. Please not modify it directly.
NAMESPACE="default"

# collector-prerequisite chart version. It will be modified in integrate building.
COLLECTOR_PREREQUISITE_CHART_VERSION=version

# subcommand, default is deploy.
# Support deploy and delete.
# This variable will not change during script executing.
SUBCOMMAND="deploy"

# ==========  end of global variables ==========


# This function used to print the help document and exit the process with code 0
function usage {
    echo "Manage kubernetes-resource-information-collector tool"
    echo "Usage:"
    echo ""
    echo "$0 <subcommand> [options...]"
    echo "subcommand:"
    echo "help                Print help documents."
    echo "version             Print version information."
    echo "deploy              Deploy kubernetes-resource-information-collector tool"
    echo "delete              Delete kubernetes-resource-information-collector tool"
    echo ""
    echo ""
    echo "options:"
    echo "-n, --namespace     Mandatory. Kubernetes namespace. The management tool will deploy/delete the kubernetes"
    echo "                    resources under this namespace."
    echo "--loglevel          Optional. Log level for this tool. Default is 3. Support 1,2,3 and 4. If given level is"
    echo "                    greater than 4, the deployment tool will treat as 4. If given level is less than 1, the"
    echo "                    management tool will treat as 1."
    echo "                    Level description: "
    echo "                    1: Print debug, info, warning and error log."
    echo "                    2: Print info, warning and error log."
    echo "                    3: Print warning and error log."
    echo "                    4: Only print error log."
    echo ""

    exit 0
}

# This function used to parse arguments
function parse_arguments {
    while test $# -gt 0; do
        case "$1" in
            help)
                usage
                ;;
            version)
                print_version
                ;;
            deploy)
                SUBCOMMAND="deploy"
                shift
                ;;
            delete)
                SUBCOMMAND="delete"
                shift
                ;;
            --loglevel)
                shift
                reset_log_level "$1"
                shift
                ;;
            --debug)
                shift
                open_debug_mode
                ;;
            -n|--namespace)
                shift
                reset_namespace "$1"
                shift
                ;;
            *)
                break
                ;;
        esac
    done
}

# Print debug log. Only print debug log while LOG_LEVEL is equal to 1
function debug {
    if [ "${LOG_LEVEL}" -ne 1 ]; then
        return
    fi

    message=$1
    if [[ "${message}" == "" ]]; then
        message="unknown message"
    fi

    message="[$(datetime)][DEBUG] ${message}"
    echo -e "${message}"
}

# Print info log. Only print info log while LOG_LEVEL is less than 3(LOG_LEVEL=1 or LOG_LEVEL=2)
function info {
    if [ "${LOG_LEVEL}" -gt 2 ]; then
        return
    fi

    message=$1
    if [[ "${message}" == "" ]]; then
        message="unknown message"
    fi

    message="[$(datetime)][INFO] ${message}"
    echo -e "${message}"
}

# Print warn log while LOG_LEVEL is less than 4
function warn {
    if [ "${LOG_LEVEL}" -eq 4 ]; then
        return
    fi

    message=$1
    if [[ "${message}" == "" ]]; then
        message="unknown message"
    fi

    message="[$(datetime)][WARNING] ${message}"
    echo -e "${message}"
}

# Print error log
function error {
    message=$1
    if [[ "${message}" == "" ]]; then
        message="unknown message"
    fi

    message="[$(datetime)][ERROR] ${message}"
    echo -e "${message}"
}

# Return a time format string that contain year, month, day, hour, minute and second. For example 2022-01-26T22:31:17
function datetime {
    echo "$(date +%F)T$(date +%T)"
}

# Open the debug mode
function open_debug_mode {
    echo "open hidden debug mode..."
    DEBUG_MODE="true"
    reset_log_level 1

    debug "---------- shell global variables before main function ----------"
    debug "DEBUG_MODE=${DEBUG_MODE}"
    debug "LOG_LEVEL=${LOG_LEVEL}"
    debug "NAMESPACE=${NAMESPACE}"
    debug "COLLECTOR_PREREQUISITE_CHART_VERSION=${COLLECTOR_PREREQUISITE_CHART_VERSION}"
    debug "SUBCOMMAND=${SUBCOMMAND}"
    debug ""
}

# Reset LOG_LEVEL. If given level is greater than 4, the function will set LOG_LEVEL to 4.
# If given level is less than 1,  function will set LOG_LEVEL to 1.
# If the value of DEBUG_MODE is true, function will set LOG_LEVEL to 1.
function reset_log_level {
    new_level=$1
    if [[ "${new_level}" == "" ]]; then
        return
    elif [ "${new_level}" -gt 4 ]; then
        new_level=4
    elif [ "${new_level}" -lt 1 ]; then
        new_level=1
    fi

    if [[ "${DEBUG_MODE}" == "true" ]]; then
        new_level=1
    fi

    LOG_LEVEL=${new_level}
}

# Reset kubernetes namespace. The function may exit with code 1 that the namespace name is not exit in kubernetes
# cluster.
function reset_namespace {
    new_namespace=$1
    if [[ "${new_namespace}" == "" ]]; then
        error "can not use a empty namespace to deploy"
        exit 1
    fi

    kubectl get ns "${new_namespace}" > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        error "error from server (NotFound): namespaces \"${new_namespace}\" not found"
        exit 1
    fi

    NAMESPACE="${new_namespace}"
}

# Get version information
function get_version {
    echo VERSION
}

# Print version information and exit with code 0
function print_version {
    version=$(get_version)
    echo "${version}"
    exit 0
}

# Deploy collector-prerequisite chart
function deploy_collector_prerequisite {
    info "install collector prerequisite..."
    work_directory=$(cd $(dirname $0) && pwd)
    helm install "$(get_version)-prerequisite" "${work_directory}/collector-prerequisite-${COLLECTOR_PREREQUISITE_CHART_VERSION}.tgz" -f "${work_directory}/values-override-collector-prerequisite.yaml" -n "${NAMESPACE}"
    if [ $? -ne 0 ]; then
        error "deploy collector-prerequisite failed..."
        exit 1
    fi
    info "deploy collector-prerequisite success."
}

# Delete collector-prerequisite chart
function delete_collector_prerequisite {
    info "delete collector prerequisite..."
    helm delete "$(get_version)-prerequisite" -n "${NAMESPACE}"
    if [ $? -ne 0 ]; then
        error "delete collector-prerequisite failed..."
        exit 1
    fi
    info "delete collector-prerequisite success."
}

# Deploy kubernetes-resource-information-collector tool
function deploy {
    deploy_collector_prerequisite
}

# Delete kubernetes-resource-information-collector tool
function delete {
    delete_collector_prerequisite
}

# The main method for this script
function main {
    if [[ "${SUBCOMMAND}" == "deploy" ]]; then
        deploy
    elif [[ "${SUBCOMMAND}" == "delete" ]]; then
        delete
    else
        error "unsupported subcommand: ${SUBCOMMAND}"
        exit 1
    fi
    exit 0
}

parse_arguments "$@"
main
