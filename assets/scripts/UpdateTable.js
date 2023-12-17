function updateTable() {
    $.getJSON('/api/v1/game/all', function (body) {
        let rows = "";

        if (body['data'] == null) {
            $('#gameTable').html(``);
            return;
        }

        body['data'].forEach(function (obj) {
            rows += `<div class="card" data-name=${obj.name}>
                        <div class="poster">
                            <input class="form-check-input" type="checkbox"
                                   value=${obj.id}
                                   id="flexCheckChecked"`

            if (obj.done) {
                rows += `checked>`;
            } else {
                rows += `>`;
            }

            rows += `<img src="../assets/static/tmpGrid.png">
                        <button type="submit" id="buttonDeleteElem"
                                class="btn btn-danger btn-sm text-center deleteElemButton" name="buttonDelete"
                                value="${obj.id}">
                        </button>
                        </div>
                        <div class="details" id="gameDetails">
                            <h1>${obj.name}</h1>
                        </div>
                    </div>`;
        });


        $('#gameTable').html(rows);
    });
}