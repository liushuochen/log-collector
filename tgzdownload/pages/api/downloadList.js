const testAPIreturn = (req, res) => {
    res.send({
        code: 200, resource: [
            {
                id: '001',
                tgzName: 'Test aaaa',
                version: 'v1.0.0',
                updateDate: '2022.04.05',
                downloadCount: 11
            },
            {
                id: '002',
                tgzName: 'Test bbbb',
                version: 'v1.0.1',
                updateDate: '2022.04.05',
                downloadCount: 12
            },
            {
                id: '003',
                tgzName: 'Test cccc',
                version: 'v1.0.2',
                updateDate: '2022.04.05',
                downloadCount: 13
            },
            {
                id: '004',
                tgzName: 'Test dddd',
                version: 'v1.0.3',
                updateDate: '2022.04.05',
                downloadCount: 14
            }
        ]
    });
}

export default testAPIreturn;