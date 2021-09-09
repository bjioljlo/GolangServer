let myChart = null;
function setTomorrowData(json) {
    console.log(json);
    document.getElementById('body').innerHTML = "";
    // Creating and adding data to first row of the table
    for (i = 1; i < (Object.keys(json).length / 6) + 1; i++) {
        let row_1 = document.createElement('tr');
        let heading_1 = document.createElement('td');
        heading_1.innerHTML = json["save_number" + i];
        let heading_2 = document.createElement('td');
        heading_2.innerHTML = json["save_number" + i + "_Date"];
        let heading_3 = document.createElement('td');
        heading_3.innerHTML = json["save_number" + i + "_UnitNumber"];
        let heading_4 = document.createElement('td');
        heading_4.innerHTML = json["save_number" + i + "_BuyPrice"];
        let heading_5 = document.createElement('td');
        heading_5.innerHTML = json["save_number" + i + "_SellPrice"];
        let heading_6 = document.createElement('td');
        heading_6.innerHTML = json["save_number" + i + "_Buy_sell"];
        let heading_7 = document.createElement('td');
        if (parseFloat(json["save_number" + i + "_BuyPrice"]) > parseFloat(json["save_number" + i + "_SellPrice"])) {
            heading_7.innerHTML = "多單"
        }
        else {
            heading_7.innerHTML = "空單"
        }

        row_1.appendChild(heading_1);
        row_1.appendChild(heading_2);
        row_1.appendChild(heading_3);
        row_1.appendChild(heading_4);
        row_1.appendChild(heading_5);
        row_1.appendChild(heading_6);
        row_1.appendChild(heading_7);

        document.getElementById('body').appendChild(row_1);
    }
}
function setReturnData(json) {
    console.log(json);
    let numbers = [];
    let returns = [];
    for (i = 0; i < Object.keys(json).length / 2; i++) {
        numbers.push(json["save_number" + (i + 1)]);
        returns.push(json["save_number" + (i + 1) + "Return"]);
    }
    feather.replace({ 'aria-hidden': 'true' })

    // Graphs
    var ctx = document.getElementById('myChart')
    if(myChart != null){
        myChart.destroy();
    }
    // eslint-disable-next-line no-unused-vars
    myChart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: numbers,
            datasets: [{
                data: returns,
                lineTension: 0,
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255,99,132,1)',
                borderWidth: 4,
                pointBackgroundColor: 'rgba(255, 99, 132, 0.2)'
            }]
        },
        options: {
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero: false
                    }
                }]
            },
            legend: {
                display: false
            }
        }
    })
}
function tomorrowData(kind) {
    $.ajax({
        //告訴程式表單要傳送到哪裡                                         
        url: "/Tdata",
        //需要傳送的資料
        data: "&kind=" + kind,
        //使用POST方法     
        type: "GET",
        //接收回傳資料的格式，在這個例子中，只要是接收true就可以了
        dataType: 'json',
        //傳送失敗則跳出失敗訊息      
        error: function () {
            //資料傳送失敗後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            alert("失敗");
        },
        //傳送成功則跳出成功訊息
        success: function (response) {
            //資料傳送成功後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            //json = JSON.parse(response.responseText)
            setTomorrowData(response)
            returnData(kind)
        }
    });
}
function returnData(kind) {
    $.ajax({
        //告訴程式表單要傳送到哪裡                                         
        url: "/Bdata",
        //需要傳送的資料
        data: "&kind=" + kind,
        //使用POST方法     
        type: "GET",
        //接收回傳資料的格式，在這個例子中，只要是接收true就可以了
        dataType: 'json',
        //傳送失敗則跳出失敗訊息      
        error: function () {
            //資料傳送失敗後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            alert("失敗");
        },
        //傳送成功則跳出成功訊息
        success: function (response) {
            //資料傳送成功後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            //json = JSON.parse(response.responseText)
            setReturnData(response)
            if (kind == "1"){
                var btn = document.getElementById("btnGroupDrop1")
                var name = document.getElementById("username")
                btn.innerHTML = name.innerText;

            }else{
                var btn = document.getElementById("btnGroupDrop1");
                btn.innerHTML = "Search排行表";
            }
        }
    });
}
function processFormData() {
    var Element = document.getElementById("stockname");
    var name = Element.value;
    $.ajax({
        //告訴程式表單要傳送到哪裡                                         
        url: "/search",
        //需要傳送的資料
        data: "&stock=" + name,
        //使用POST方法     
        type: "GET",
        //接收回傳資料的格式，在這個例子中，只要是接收true就可以了
        dataType: 'json',
        //傳送失敗則跳出失敗訊息      
        error: function () {
            //資料傳送失敗後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            alert("失敗");
        },
        //傳送成功則跳出成功訊息
        success: function () {
            //資料傳送成功後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            //alert("成功");
            location.href = '/stock?stock=' + name;
        }
    });
}
function saveTtock() {
    var Element = document.getElementById("stockname");
    var name = Element.value;
    $.ajax({
        //告訴程式表單要傳送到哪裡                                         
        url: "/save",
        //需要傳送的資料
        data: "&stock=" + name,
        //使用POST方法     
        type: "GET",
        //接收回傳資料的格式，在這個例子中，只要是接收true就可以了
        dataType: 'json',
        //傳送失敗則跳出失敗訊息      
        error: function () {
            //資料傳送失敗後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            alert("失敗");
        },
        //傳送成功則跳出成功訊息
        success: function () {
            //資料傳送成功後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            //alert("成功");
            location.reload();
        }
    });
}
function deletTtock() {
    var Element = document.getElementById("stockname");
    var name = Element.value;
    $.ajax({
        //告訴程式表單要傳送到哪裡                                         
        url: "/delet",
        //需要傳送的資料
        data: "&stock=" + name,
        //使用POST方法     
        type: "GET",
        //接收回傳資料的格式，在這個例子中，只要是接收true就可以了
        dataType: 'json',
        //傳送失敗則跳出失敗訊息      
        error: function () {
            //資料傳送失敗後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            alert("失敗");
        },
        //傳送成功則跳出成功訊息
        success: function () {
            //資料傳送成功後就會執行這個function內的程式，可以在這裡寫入要執行的程式  
            //alert("成功");
            location.reload();
        }
    });
}
$('.loader-inner').loaders();