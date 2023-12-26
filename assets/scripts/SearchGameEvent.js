$(document).ready(function () {
    $("#search").keyup(function () {
        _this = this;
        $.each($("#gameTable div.card"), function () {
            if (this.getAttribute("data-name").toLowerCase().indexOf($(_this).val().toLowerCase()) == -1) {
                $(this).hide();
            } else {
                $(this).show();
            }
        });
    });
});