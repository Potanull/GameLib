$(document).on('click', '#buttonRandom', function (event) {
    $.ajax({
        url: '/api/v1/game/random',
        method: 'GET',
        dataType: 'json',
        success: function (body) {
            const element = $("#labelRandomGame");
            const newText = body.data.name;
            element.text(newText);
            $("#randomGame").show();
            setTimeout(function () {
                $('#randomGame').hide();
            }, 8000);
        }
    });
});
