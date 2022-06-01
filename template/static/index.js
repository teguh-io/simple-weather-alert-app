dataWindObj = document.getElementById('data-wind').innerHTML
dataWind = parseInt(dataWindObj)
windStatus = ""

if (dataWind <= 6) {
    windStatus = "AMAN"
    document.getElementById('wind-status').innerHTML = windStatus    
}

else if (dataWind >= 7 && dataWind <= 15) {
    windStatus = "SIAGA"
    document.getElementById('wind-status').innerHTML = windStatus
    document.getElementById('wind').className = "background-animation-siaga"
}
else if (dataWind > 15 ) {
    windStatus = "BAHAYA"
    document.getElementById('wind-status').innerHTML = windStatus
    document.getElementById('wind').className = "background-animation-bahaya"
}


// Water
dataWaterObj = document.getElementById('data-water').innerHTML
dataWater = parseInt(dataWaterObj)
waterStatus = ""

if (dataWater <= 5) {
    waterStatus = "AMAN"
    document.getElementById('water-status').innerHTML = waterStatus    
}

else if (dataWater >= 6 && dataWater <= 8) {
    waterStatus = "SIAGA"
    document.getElementById('water-status').innerHTML = waterStatus
    document.getElementById('water').className = "background-animation-siaga"
}
else if (dataWater > 8 ) {
    waterStatus = "BAHAYA"
    document.getElementById('water-status').innerHTML = waterStatus
    document.getElementById('water').className = "background-animation-bahaya"
}

// Overall
if ((windStatus == "AMAN") && (waterStatus == "AMAN")) {
    document.getElementById('overall-status').innerHTML = "AMAN"
    document.getElementById('circle').style.outline = "#4CAF50 solid 10px"
}
else if ((windStatus == "AMAN") && (waterStatus == "SIAGA")) {
    document.getElementById('overall-status').innerHTML = "SIAGA BANJIR"   
    document.getElementById('circle').className = "circle-animation-siaga"
}
else if ((windStatus == "AMAN" || windStatus == "SIAGA") && (waterStatus == "BAHAYA")) {
    console.log("BAHAYA BANJIR")   
    document.getElementById('overall-status').innerHTML = "BAHAYA BANJIR" 
    document.getElementById('circle').className = "circle-animation-bahaya"   
}
else if ((windStatus == "SIAGA") && (waterStatus == "AMAN")) {
    document.getElementById('overall-status').innerHTML = "SIAGA BADAI"  
    document.getElementById('circle').className = "circle-animation-siaga"  
}

else if ((waterStatus == "AMAN" || waterStatus == "SIAGA") && (windStatus == "BAHAYA")) {
    document.getElementById('overall-status').innerHTML = "BAHAYA BADAI"
    document.getElementById('circle').className = "circle-animation-bahaya"
}
else if ((windStatus == "BAHAYA") && (waterStatus == "BAHAYA")) {
    document.getElementById('overall-status').innerHTML = "BAHAYA TOPAN"
    document.getElementById('circle').className = "circle-animation-bahaya"
}

else if ((windStatus == "SIAGA") && (waterStatus == "SIAGA")) {
    document.getElementById('overall-status').innerHTML = "SIAGA TOPAN"
    document.getElementById('circle').className = "circle-animation-siaga"
}