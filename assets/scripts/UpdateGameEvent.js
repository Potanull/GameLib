function displaySelectedImage(event, elementId) {
    const selectedImage = document.getElementById(elementId);
    const fileInput = event.target;

    if (fileInput.files && fileInput.files[0]) {
        const reader = new FileReader();

        reader.onload = function (e) {
            selectedImage.src = e.target.result;
        };

        reader.readAsDataURL(fileInput.files[0]);
    }
}

$(document).on('click', '#buttonUpdateElem', function (event) {
    let id = $(this).val()
    $.getJSON('/api/v1/game/' + id, function (body) {
        let rows = `
            <div class="col-12 mb-2">
              <label for="updateGameNameInput" class="form-label">Название</label>
              <input type="text" class="form-control" name="game" id="updateGameNameInput" 
                        autocomplete="off" value="${body.data.name}" placeholder="Введите название игры...">
            </div>
            <div class="col-12 form-check mb-2">
                <input type="checkbox" class="form-check-input" id="updateGameStatusInput"`;

        if (body.data.done) {
            rows += `checked>`;
        } else {
            rows += `>`;
        }

        rows += `<label class="form-check-label" for="updateGameStatusInput">Статус</label>
            </div>   
            <div class="inputGrid mb-2">
                <div class="mb-4 d-flex justify-content-center">
                    <img id="selectedImage" src="${body.data.image_url}" class="mx-auto d-block">
                </div>
            </div>
            <label class="btn btn-primary btn-rounded form-label text-white m-1" for="updateGridButton">Обложка</label>
            <input type="file" class="form-control d-none" accept="image/png, image/jpeg" id="updateGridButton" onchange="displaySelectedImage(event, 'selectedImage')"/>
        `;

        $('#updateGameBody').html(rows);
        $('#buttonUpdate').attr('value', id);
    });
});

function getInfoForUpdate() {
    let obj;

    if ($('#updateGridButton').get(0).files[0]) {
        obj = {
            name: $("#updateGameNameInput").val(),
            done: $("#updateGameStatusInput").is(":checked"),
            image_url: "../assets/static/grids/" + $('#updateGridButton').get(0).files[0].name
        };
    } else {
        obj = {
            name: $("#updateGameNameInput").val(),
            done: $("#updateGameStatusInput").is(":checked"),
        };
    }

    return JSON.stringify(obj);
}

function updateGame(id) {
    $.ajax({
        url: '/api/v1/game/' + id,
        method: 'PUT',
        dataType: 'json',
        data: getInfoForUpdate(),
        statusCode: {
            200: function () {
                alert("Игра уже есть в списке");
            },
            201: async function () {
                await saveImg($('#updateGridButton').get(0).files[0], $('#updateGridButton').get(0).files[0].name);
                $('#updateGameModal').modal('toggle');
                updateTable();
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

$(document).on('click', '#buttonUpdate', function (event) {
    updateGame($(this).val())
});

$('#updateGameNameInput').bind("enterKey", function (e) {
    updateGame($(this).val())
});

$('#updateGameNameInput').keyup(function (e) {
    if (e.keyCode == 13) {
        $(this).trigger("enterKey");
    }
});