import React from 'react'
import styles from '../styles/download.module.css'
import _ from 'lodash'

class downloadTable extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            tgzResource: this.props.posts.resource,
        }
    }

    render() {
        const CreateRow = (props) => {
            const items = props.items;
            const listItem = items.map(item =>
                <tr key={item.id}>
                    <td>{item.tgzName}</td>
                    <td>{item.version}</td>
                    <td>{item.updateDate}</td>
                    <td>{item.downloadCount}</td>
                </tr>
            )

            return (
                <table className={styles.downloadTable}>
                    <thead>
                        <tr>
                            <td>tgz资源包名称</td>
                            <td>版本号</td>
                            <td>最后更新日期</td>
                            <td>下载次数</td>
                        </tr>
                    </thead>
                    <tbody>
                        {listItem}
                    </tbody>
                </table>
            )
        }

        return (
            <div>
                <header className={styles.pageHeader}>
                    Welcome to KUBERNETES-RESOURCE-INFORMATION-COLLECTOR
                </header>
                <CreateRow items={this.state.tgzResource} />
            </div>
        )
    }
}

export const getStaticProps = async () => {
    const res = await fetch('http://localhost:3000/api/downloadList', {
        method: 'get',
    });
    const posts = await res.json();
    // 通过返回 { props: { posts } } 对象，上面组件在构建时将接收到 `posts` 参数
    return {
        props: {
            posts,
        },
    }
}

export default downloadTable;