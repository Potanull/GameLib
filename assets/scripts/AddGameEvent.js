function getInfo() {
    let obj = {
        name: $("#inputNewGame").val(),
        image: $('#gridForGame').get(0).files[0].name
    };
    return JSON.stringify(obj);
}

function aggregationImg() {
    var formData = new FormData();
    formData.append('image', $('#gridForGame').get(0).files[0]);
    return formData
}

async function saveImg() {
    $.ajax({
        url: 'api/image/' + $('#gridForGame').get(0).files[0].name,
        method: 'POST',
        data: aggregationImg(),
        async: false,
        contentType: false, // NEEDED, DON'T OMIT THIS (requires jQuery 1.6+)
        processData: false, // NEEDED, DON'T OMIT THIS
    });
}

function createGame() {
    $.ajax({
        url: '/api/v1/game/',
        method: 'POST',
        dataType: 'json',
        data: getInfo(),
        statusCode: {
            201: async function () {
                saveImg();
                $('#inputNewGame').val("");
                $('#gridForGame').val("");
                updateTable();
                $('#addGameModal').modal('toggle');
            },
            200: function () {
                alert("Игра уже есть в списке");
            },
            422: function () {
                alert("Имя игры слишком большое или пустое");
            },
            400: function () {
                alert("Что-то пошло не так!");
            }
        }
    });
}

$(document).on('click', '#buttonAdd', function (event) {
    createGame()
});

$('#inputNewGame').bind("enterKey", function (e) {
    createGame()
});

$('#inputNewGame').keyup(function (e) {
    if (e.keyCode == 13) {
        $(this).trigger("enterKey");
    }
});