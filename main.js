const fs = require('fs');
const dailyList = require('./data.json');

const dailyCases = [];

let length = dailyList.cases_time_series.length;

for(let i = dailyList.cases_time_series.length; i--;){
      dailyCases.push([dailyList.cases_time_series[i].date, dailyList.cases_time_series[i].dailyconfirmed]);
}

fs.writeFileSync('nodeCovid.csv', JSON.stringify(dailyCases));