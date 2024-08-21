

function changeBayMonthly(newTitle,bay) {
    document.getElementById('dropdownTitle').textContent = newTitle;
    console.log('Hello',bay)
   monthlyChangeBay(bay)
   //dailyChangeBay(bay)
}
function changeBayDaily(newTitle,bay) {
    document.getElementById('dropdownTitle').textContent = newTitle;
    console.log('Hello',bay)
   // monthlyChangeBay(bay)
   dailyChangeBay(bay)
}

function changeBayYearly(newTitle,bay) {
    document.getElementById('dropdownTitle').textContent = newTitle;
   yearlyChangeBay(bay)
}

function dailyChangeBay(bay){
    htmx.ajax('GET',`/daily-data?bay=${bay}`,'#body-table').then(()=>{
    });
   // hx-get="/daily-data" hx-target="#body-table" hx-trigger="click" hx-vals='{"bay":1}'
}
function yearlyChangeBay(bay){
    // htmx.ajax('GET',`/yearly-light?bay=${bay}`,'#yearly-light').then(()=>{
    // });
    htmx.ajax('GET',`/yearly-light?bay=${bay}`,'#yearly-light').then(()=>{
    });
    htmx.ajax('GET',`/yearly-peak?bay=${bay}`,'#yearly-peak').then(()=>{});
   // hx-get="/daily-data" hx-target="#body-table" hx-trigger="click" hx-vals='{"bay":1}'
}

function monthlyChangeBay(bay){
    htmx.ajax('GET',`/day-time-peak?bay=${bay}`,'#body-table-day').then(()=>{
        htmx.ajax('GET',`/night-time-peak?bay=${bay}`,'#body-table-night').then(()=>{
        console.log(`get content night-time-peak bay = ${bay}`)
    });
    htmx.ajax('GET',`/all-time-peak?bay=${bay}`,'#body-table-all').then(()=>{
        console.log('get content all-time-peak')
    });
    });
   
}