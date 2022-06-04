import _ from 'lodash';

const getDownloadList = async (req) => {
    let result = {};
    _.set(result, 'data', 'test user API route');
    return result;
};

export default {
    getDownloadList,
}