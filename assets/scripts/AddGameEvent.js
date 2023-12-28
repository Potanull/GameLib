function getInfoForAdd() {
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
        data: getInfoForAdd(),
        statusCode: {
            200: function () {
                alert("Игра уже есть в списке");
            },
            201: async function () {
                await saveImg($('#gridForGame').get(0).files[0], $('#gridForGame').get(0).files[0].name);
                clearInputForm();
                updateTable();
                $('#addGameModal').modal('toggle');
            },
            400: function () {
                alert("Что-то пошло не так!");
            },
            422: function () {
                alert("Имя игры слишком большое или пустое");
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