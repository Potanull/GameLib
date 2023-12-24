$(document).on('click', '#buttonAdd', function (event) {
    $.ajax({
        url: '/api/v1/game/',
        method: 'POST',
        dataType: 'json',
        data: JSON.stringify({name: $("#inputNewGame").val()}),
        statusCode: {
            201: function () {
                $('#inputNewGame').val("");
                updateTable();
            },
            200: function () {
                alert("Игра уже есть в списке");
            },
            422: function () {
                alert("Имя игры слишком большое");
            },
            400: function () {
                alert("Что-то пошло не так!");
            }
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            alert("Что-то пошло не так!");
        }
    });
});

$('#inputNewGame').bind("enterKey", function (e) {
    $.ajax({
        url: '/api/v1/game/',
        method: 'POST',
        dataType: 'json',
        data: JSON.stringify({name: $("#inputNewGame").val()}),
        success: function () {
            $('#inputNewGame').val("");
            updateTable();
        },
        error: function (XMLHttpRequest, textStatus, errorThrown) {
            alert("Игра уже есть в списке");
        }
    });
});

$('#inputNewGame').keyup(function (e) {
    if (e.keyCode == 13) {
        $(this).trigger("enterKey");
    }
});