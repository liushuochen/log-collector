// import mongoose module
import mongoose from 'mongoose';
// connect database serve
mongoose.connect('mongodb://localhost:27017/localhostDB', {
    useNewUrlParser: true,
    useUnifiedTopology: true
}, function (error) {
    if (error) {
        console.log("数据库连接失败")
    } else {
        console.log("数据库连接成功")
    }
})
// export
export default mongoose;
