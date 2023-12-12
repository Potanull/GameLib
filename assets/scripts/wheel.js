function getListGames() {
    let res = [];

    $.ajax({
        url: '/api/v1/game/random/list',
        method: 'GET',
        dataType: 'json',
        async: false,
        success: function (body) {
            let games = body.data;
            for (let i = 0; i < games.length; i++) {
                res.push({
                    label: games[i].name,
                    value: games[i].name,
                });
            }
        }
    });

    return res;
}

$(document).on('click', '#buttonSpin', function () {
    $("#chart").html("");
    $(".rollGame").html("ㅤ");

    var padding = {top: 20, right: 20, bottom: 0, left: 20},
        w = 750 - padding.left - padding.right,
        h = 750 - padding.top - padding.bottom,
        r = Math.min(w, h) / 2,
        rotation = 0,
        oldrotation = 0,
        picked = 100000,
        oldpick = [],
        color = ["#0d6efd", "#0445a4"]


    var data = getListGames();

    var svg = d3.select('#chart')
        .append("svg")
        .data([data])
        .attr("width", w + padding.left + padding.right)
        .attr("height", h + padding.top + padding.bottom);

    var container = svg.append("g")
        .attr("class", "chartholder")
        .attr("transform", "translate(" + (w / 2 + padding.left) + "," + (h / 2 + padding.top) + ")");

    var vis = container
        .append("g");

    var pie = d3.layout.pie().sort(null).value(function (d) {
        return 1;
    });

// declare an arc generator function
    var arc = d3.svg.arc().outerRadius(r);

// select paths, use arc generator to draw
    var arcs = vis.selectAll("g.slice")
        .data(pie)
        .enter()
        .append("g")
        .attr("class", "slice");


    arcs.append("path")
        .attr("fill", function (d, i) {
            return color[i % 2];
        })
        .attr("d", function (d) {
            return arc(d);
        });

    // add the text
    arcs.append("text").attr("transform", function (d) {
        d.innerRadius = 0;
        d.outerRadius = r;
        d.angle = (d.startAngle + d.endAngle) / 2;
        return "rotate(" + (d.angle * 180 / Math.PI - 90) + ")translate(" + (d.outerRadius - 10) + ")";
    }).attr("text-anchor", "end").style({"fill": "white"}).text(function (d, i) {
            return data[i].label;
    });

    container.on("click", spin);

    //make arrow
    // svg.append("g")
    //     .attr("transform", "translate(" + (w + padding.left + padding.right) + "," + ((h / 2) + padding.top) + ")")
    //     .append("path")
    //     .attr("d", "M-" + (r * .15) + ",0L0," + (r * .05) + "L0,-" + (r * .05) + "Z")
    //     .style({"fill": "black", "margin": "auto",});

    //draw spin circle
    container.append("circle")
        .attr("cx", 0)
        .attr("cy", 0)
        .attr("r", 30)
        .style({"fill": "white", "cursor": "pointer"});

    //spin text
    container.append("text")
        .attr("x", 0)
        .attr("y", 0)
        .attr("text-anchor", "middle")
        .text("")
        .style({"font-weight": "bold", "font-size": "30px", "fill": "black"});

    function spin(d) {

        container.on("click", null);

        //all slices have been seen, all done
        if (oldpick.length == data.length) {
            container.on("click", null);
            return;
        }

        var ps = 360 / data.length,
            pieslice = Math.round(1440 / data.length),
            rng = Math.floor((Math.random() * 1440) + 360);

        rotation = (Math.round(rng / ps) * ps);

        picked = Math.round(data.length - (rotation % 360) / ps);
        picked = picked >= data.length ? (picked % data.length) : picked;


        if (oldpick.indexOf(picked) !== -1) {
            d3.select(this).call(spin);
            return;
        } else {
            oldpick.push(picked);
        }

        rotation += 90 - Math.round(ps / 2);

        vis.transition()
            .duration(3000)
            .attrTween("transform", rotTween)
            .each("end", function () {
                oldrotation = rotation;

                d3.select(".slice:nth-child(" + (picked + 1) + ") path").attr("fill", "#212529");

                d3.select("#value h1").style({"color": "black","font-size": "20px"}).text(data[picked].value);

                container.on("click", spin);
            });
    }

    function rotTween(to) {
        var i = d3.interpolate(oldrotation % 360, rotation);
        return function (t) {
            return "rotate(" + i(t) + ")";
        };
    }


    function getRandomNumbers() {
        var array = new Uint16Array(1000);
        var scale = d3.scale.linear().range([360, 1440]).domain([0, 100000]);

        if (window.hasOwnProperty("crypto") && typeof window.crypto.getRandomValues === "function") {
            window.crypto.getRandomValues(array);
        } else {
            //no support for crypto, get crappy random numbers
            for (var i = 0; i < 1000; i++) {
                array[i] = Math.floor(Math.random() * 100000) + 1;
            }
        }

        return array;
    }
})