function getInfo() {
    let obj;

    if ($('#gridForGame').get(0).files[0]) {
        obj = {
            name: $("#inputNewGame").val(),
            done: $(".flexChecked").is(":checked"),
            image: $('#gridForGame').get(0).files[0].name
        };
    } else {
        obj = {
            name: $("#inputNewGame").val(),
            done: $(".flexChecked").is(":checked"),
        };
    }

    return JSON.stringify(obj);
}

function aggregationImg() {
    var formData = new FormData();
    formData.append('image', $('#gridForGame').get(0).files[0]);
    return formData
}

async function saveImg() {
    if ($('#gridForGame').get(0).files[0]) {
        $.ajax({
            url: 'api/image/' + $('#gridForGame').get(0).files[0].name,
            method: 'POST',
            data: aggregationImg(),
            async: false,
            contentType: false, // NEEDED, DON'T OMIT THIS (requires jQuery 1.6+)
            processData: false, // NEEDED, DON'T OMIT THIS
        });
    }
}

function clearInputForm() {
    $('#inputNewGame').val("");
    $('#gridForGame').val("");
    $(".flexChecked").prop('checked', false);
}

function createGame() {
    $.ajax({
        url: '/api/v1/game/',
        method: 'POST',
        dataType: 'json',
        data: getInfo(),
        statusCode: {
            201: async function () {
                await saveImg();
                clearInputForm();
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