function displaySelectedImage(event, elementId) {
    const selectedImage = document.getElementById(elementId);
    const fileInput = event.target;

    if (fileInput.files && fileInput.files[0]) {
        const reader = new FileReader();

        reader.onload = function(e) {
            selectedImage.src = e.target.result;
        };

        reader.readAsDataURL(fileInput.files[0]);
    }
}

$(document).on('click', '#buttonUpdateElem', function (event) {
    $.getJSON('/api/v1/game/' + $(this).val(), function (body) {
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
            <input type="file" class="form-control d-none" id="updateGridButton" onchange="displaySelectedImage(event, 'selectedImage')"/>
        `;

        $('#updateGameBody').html(rows);
    });
});