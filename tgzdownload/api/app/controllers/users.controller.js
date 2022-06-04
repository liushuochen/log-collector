import usersHandler from '../handlers/users.handlers.js';

// get tgzDownloadList from DB
const getDownloadList = async (req, res) => {
    let result = {};
    result = await usersHandler.getDownloadList(req);
    res.status(200).send(result);
};

export default {
    getDownloadList,
}