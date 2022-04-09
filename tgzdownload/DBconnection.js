const mysql = require('mysql')
const CONFIG = require('./config.json')

const connection = mysql.createConnection({
  host: CONFIG.mysqlOptions.host,
  user: CONFIG.mysqlOptions.username,
  password: CONFIG.mysqlOptions.password,
  database: CONFIG.mysqlOptions.database
});

connection.connect((err) => {
  if (err) {
    console.log(err);
  }
});


connection.query('SELECT 1 + 1 AS solution', function (error, results, fields) {
  if (error) throw error;
  console.log('The solution is: ', results[0].solution);
});