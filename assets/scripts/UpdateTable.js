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
                            <input class="form-check-input statusCheckBox" type="checkbox"
                                   value=${obj.id} `

            if (obj.done) {
                rows += `checked>`;
            } else {
                rows += `>`;
            }

            if (obj.image_url) {
                rows += `<img src="${obj.image_url}">`
            } else {
                rows += `<img src="../assets/static/tmpGrid.png">`
            }
            rows += `<button type="button" id="buttonDeleteElem"
                            class="btn btn-danger btn-sm text-center deleteElemButton" name="buttonDelete"
                            value="${obj.id}">
                    </button>
                    <button type="button" id="buttonUpdateElem" data-bs-toggle="modal" data-bs-target="#updateGameModal"
                            class="btn btn-primary btn-sm text-center updateElemButton" name="buttonUpdate"
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