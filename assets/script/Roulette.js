$('#buttonRoulette').on('click', function() {
    setRoulette();
});

$('#buttonUpdateRoulette').on('click', function(){
    $('.roulette-wrapper .roulette').html("")
    setRoulette();
});

$('#buttonSpinRoulette').on('click', function(){
    spinRoulette();
});

function setRoulette() {
    let randomGameList = []
    $.ajax({
        url: '/api/v1/game/random/list?image=true',
        method: 'GET',
        dataType: 'json',
        async: false,
        success: function (body) {
            let games = body.data;
            for (let i = 0; i < games.length; i++) {
                randomGameList.push({
                    name: games[i].name,
                    image_url: games[i].image_url,
                });
            }
        }
    });

    initRoulette(randomGameList);
}

function initRoulette(games){
    let $wheel = $('.roulette-wrapper .roulette'),
        row = "";

    row += "<div class='rowRoulette'>";
    for (let i = 0; i < games.length; i++) {
        row += `<div class='card'>
                    <img src="${games[i].image_url}" alt="${games[i].name}">
                    <p>${games[i].name}</p>
                <\/div>`
    }
    row += "<\/div>";

    for(let x = 0; x < games.length; x++){
        $wheel.append(row);
    }
}

function spinRoulette(){
    let min = 0;
    let max = 29;
    let $wheel = $('.roulette-wrapper .roulette'),
        position = Math.random() * (max - min) + min;

    let rows = 12,
        card = 280 + 3 * 2,
        landingPosition = (rows * 15 * card) + (position * card);

    let randomize = Math.floor(Math.random() * 280) - (280/2);

    landingPosition = landingPosition + randomize;

    let object = {
        x: Math.floor(Math.random() * 50) / 100,
        y: Math.floor(Math.random() * 20) / 100
    };

    $wheel.css({
        'transition-timing-function':'cubic-bezier(0,'+ object.x +','+ object.y + ',1)',
        'transition-duration':'6s',
        'transform':'translate3d(-'+landingPosition+'px, 0px, 0px)'
    });

    setTimeout(function(){
        $wheel.css({
            'transition-timing-function':'',
            'transition-duration':'',
        });

        let resetTo = -(position * card + randomize);
        $wheel.css('transform', 'translate3d('+resetTo+'px, 0px, 0px)');
    }, 6 * 1000);
}