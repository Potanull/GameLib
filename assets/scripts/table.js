$(document).ready(function () {
    $('#gameTable').DataTable({
        "language": {
            "search": "Поиск:",
            "lengthMenu": "Кол-во игр на странице _MENU_",
            "loadingRecords": "Загрузка...",
            "processing": "",
            "zeroRecords": "Не найдено совпадающих игр",
            "paginate": {
                "first": "Первый",
                "last": "Последний",
                "next": "Следующий",
                "previous": "Предыдущий"
            },
        },
        ordering: false,
        info: false,
        iDisplayLength: -1,
        lengthMenu: [
            [10, 50, 100, -1],
            [10, 50, 100, 'Все'],
        ],
    });
});
