$(document).on('click', '#buttonDeleteElem', function (event) {
    $.ajax({
        url: '/api/delete',
        method: 'DELETE',
        dataType: 'json',
        data: {'deleteGame': $(this).val()},
        success: function () {
            updateTable();
        }
    });
});