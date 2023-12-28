function aggregationImg(date) {
    var formData = new FormData();
    formData.append('image', date);
    return formData
}

async function saveImg(date) {
    if (date) {
        $.ajax({
            url: 'api/image/' + date.name,
            method: 'POST',
            data: aggregationImg(date),
            async: false,
            contentType: false, // NEEDED, DON'T OMIT THIS (requires jQuery 1.6+)
            processData: false, // NEEDED, DON'T OMIT THIS
        });
    }
}