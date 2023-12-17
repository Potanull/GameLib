$(document).on('change', ':checkbox', function (event) {
    $.ajax({
        url: '/api/v1/game/reverse/status/' + $(this).val(),
        method: 'PUT',
        dataType: 'json',
        success: function () {
            updateTable();
        }
    });
});
