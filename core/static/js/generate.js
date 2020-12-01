var index = 1;

function add() {
    document.getElementById("param").innerHTML("参数"+ index + ":<input type='text' id='par" + index +"' value='name' /> <br>")
    index++
}

function remove() {
    document.getElementById("par" + index).remove()
    index--
}