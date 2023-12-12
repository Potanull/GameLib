function  updateTable() {
    $.getJSON('/api/v1/game/all', function (body) {
        let rows = `<thead>
                    <tr>
                        <th scope="col"><img src="https://img.icons8.com/ios-glyphs/24/FFFFFF/checked-2--v1.png"/></th>
                        <th scope="col" class="w-100"><img src="https://img.icons8.com/ios-glyphs/24/FFFFFF/xbox-controller.png"/>
                            Игры
                        </th>
                        <th scope="col">
                            Удалить
                        </th>
                    </tr>
                    </thead>
                    <tbody>`;

        body['data'].forEach(function (obj) {
            let row = `<tr><td class="align-middle"><input class="form-check-input" type="checkbox" value="${obj.id}" id="flexCheckChecked" `;
            if (obj.done) {
                row += `checked=""`
            }
            row += `><td class="align-middle">${obj.name}</td>`;
            row += `<td class="align-middle">
                <div class="d-flex justify-content-center">
                    <button type="submit" id="buttonDeleteElem"
                            class="btn btn-danger btn-sm text-center deleteElemButton" name="buttonDelete"
                            value="${obj.id}">
                    </button>
                </div>
            </td>`;
            row += `</tr>`;
            rows += row;
        });

        rows += `</tbody>`;

        $('#gameTable').html(rows);
    });
}

