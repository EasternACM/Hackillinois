var Bushele = {
  fieldList: null,

  field: function(IDnumber, latitude, longitude, rain, temperature, pests, soil, cropType){
    //constructor for a field object
    this.idnumber = IDnumber,
    this.gps = {lat: latitude, long: longitude},
    this.rainfall = rain,
    this.temperature = temperature,
    this.pests = pests,
    this.soilQuality = soil,
    this.crop = cropType
  },

  labelCharts: function(){
    for(i = 1; i < 5; i++){
      var title = document.getElementById(i.toString() + ' Label');
      var sub = document.getElementById(i.toString() + ' Span');

      title.innerHTML = 'Title Name';
      sub.innerHTML = 'Address';
    }
  },

  getDataFromAPI: function(){
    //call api and initialize object list
    //var obj;
    $.getJSON('http://whateverorigin.org/get?url=' + encodeURIComponent('http://zacc.xyz:8000/fields') + '&callback=?', function(data){
	  //console.log(data.contents);
    //var obj = $.parseJSON(data.contents);
    console.log(data);
    var rain = Bushele.getRain();
    var temp = Bushele.getTemperature();
    var pest = Bushele.hasPests();
    var soil = Bushele.getSoilQuality();
    var cropType = Bushele.getCropType();

    Bushele.fieldList = new Bushele.field(obj.id, obj.centroid.coordinates[0], obj.centroid.coordinates[1], rain, temp, pests, soil, cropType);
  });
  },

  getRain: function(){ //placeholder for future Climate Corporation API data
    return Math.random() * 3; //cm rainfall
  },

  getTemperature: function(){//placeholder for future Climate Corporation API data
    return (~~Math.random() * 100); //deg F
  },

  hasPests: function(){//placeholder for future Climate Corporation API data
    return (Math.round(Math.random())); //T / F
  },

  getSoilQuality: function(){//placeholder for future Climate Corporation API data
    return ((Math.random) * (50) + 0);
  },

  getCropType: function(){//placeholder for future Climate Corporation API data
    var i = ~~(Math.random() * 3);
    var strings = ["Corn", "Wheat", "Soy Beans"];
    return strings[i];
  },

  sortFields: function(modeStr){
    //ObjectArray[index][keyString]
    for (i = 1; i < fieldList.length; i++){
      j = i;
      while( j > 0 && fieldList[j-1][modeStr] > fieldList[j][modeStr]){
        //swap j and j-1
        temp = fieldList[j][modeStr];
        fieldList[j][modeStr] = fieldList[j-1][modeStr];
        fieldList[j-1][modeStr] = temp;

        j = j-1;
      }
    }
    //console.log(list);
  },

  addFields: function(number){ //modes:
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
      circleDiv.id = i.toString() + "circle";
      circleDiv.style = "border-radius:50%; background-color: #E9014C; width:40px; height:40px;";

      var title = document.createElement('H4');
      title.id = i.toString() + "title";
      title.innerHTML = 'Title Name';

      var sub = document.createElement('SPAN');
      sub.id = i.toString() + "sub";
      sub.innerHTML = 'Address';

      padding.appendChild(circleDiv);

      container.appendChild(padding);
      container.appendChild(title);
      container.appendChild(sub);

      parentDiv.appendChild(container);
    }
  },


  createCORSRequest: function(method, url){
    var xhr = new XMLHttpRequest();
    if ("withCredentials" in xhr){
        xhr.open(method, url, true);
    } else if (typeof XDomainRequest != "undefined"){
        xhr = new XDomainRequest();
        xhr.open(method, url);
    } else {
        xhr = null;
    }
    return xhr;
  },

  getDummyData: function(){
    return;
  }
};
