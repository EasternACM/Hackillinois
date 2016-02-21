function labelCharts(){
  for(i = 1; i < 5; i++){
    var title = document.getElementById(i.toString() + ' Label');
    var sub = document.getElementById(i.toString() + ' Span');

    title.innerHTML = 'Title Name';
    sub.innerHTML = 'Address';
  }
}

function GetDataFromAPI(){
  
}

function sortFields(modeStr){
  //ObjectArray[index][keyString]
  for (i = 1; i < list.length; i++){
    j = i;
    while( j > 0 && list[j-1][modeStr] > list[j][modeStr]){
      //swap j and j-1
      temp = list[j][modeStr];
      list[j][modeStr] = list[j-1][modeStr];
      list[j-1][modeStr] = temp;

      j = j-1;
    }
  }
  //console.log(list);
}

function addFields(number){ //modes:
  if(number <= 0){
    alert('Error in addFields function');
    return;
  }

  var parentDiv = document.getElementById('fieldContainer');

  for(i = 0; i < number; i++){
    var container = document.createElement('DIV');
    container.className = 'col-xs-6 col-sm-3 placeholder';

    var padding = document.createElement('DIV');
    padding.style = 'padding-left: 10px;';

    var circleDiv = document.createElement('DIV');
    circleDiv.style = "border-radius:50%; background-color:green; width:40px; height:40px;";

    var title = document.createElement('H4');
    title.innerHTML = 'Title Name';

    var sub = document.createElement('SPAN');
    sub.innerHTML = 'Address';

    padding.appendChild(circleDiv);

    container.appendChild(padding);
    container.appendChild(title);
    container.appendChild(sub);

    parentDiv.appendChild(container);
  }
}
