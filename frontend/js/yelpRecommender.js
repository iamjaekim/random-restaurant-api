$(document).ready(() => {
    $('#searchForm').on('submit keyup touchend input', (e) => {
        let zipCode = $('#zipcode').val();
        getRestaurants(zipCode);
        e.preventDefault();
    });
});

function getRestaurants(zipCode) {
    $.ajax({
        type: "GET",
        url: 'http://127.0.0.1:5000/api/stores/'+zipCode,
        contentType: "application/x-www-form-urlencoded",
    }).then((response) => {
        console.log(response)
        let stores = response;
        let output = '';
        $.each(stores, (index, store) => {
            output += `
                <div class = "col-md-3">
                    <div class="well text-center">
                        <h5>${store.name}</h5>
                        <a onclick="restaurantSelected('${store.id}')" class="btn btn-primary" href="#">Restaurant Details</a>
                    </div>
                </div>
            `
        })
        $('#restaurants').html(output)
    }).catch((err)=> {
        console.log(err)
        let output = `
            <div class="well col-md-3">
                ${err.responseText}
            </div>
        `;
        $('#restaurants').html(output);
    });
}

function restaurantSelected(storeId) {
    sessionStorage.setItem('storeId', storeId);
    window.location='restaurant.html';
    return false;
}

function getRestaurant() {
    let restaurantId = sessionStorage.getItem('storeId');

    $.ajax({
        type: "GET",
        url: 'http://127.0.0.1:5000/api/stores/single/'+restaurantId,
        contentType: "application/x-www-form-urlencoded",
    }).then((response) => {
        console.log(response)
        let restaurant = response
        let output = `
            <div class="row">
                <div class="col-md-4">
                    <img src="${restaurant.image_url}" class="thumbnail fit-image">
                </div>
                <div class="col-md-8">
                    <h2>${restaurant.name}</h2>
                    <ul class="list-group">
                        <li class="list-group-item">Address: ${restaurant.location.display_address}</li>
                        <li class="list-group-item">Phone Number: ${restaurant.phone}</li>
                        <li class="list-group-item">Price Range: ${restaurant.price}</li>
                        <li class="list-group-item">Rating: ${restaurant.rating}</li>
                        <li class="list-group-item">Number of Reviews: ${restaurant.review_count}</li>
                        <li class="list-group-item">            
                        <div class="row">
                          <div class="well">
                            <a href="${restaurant.url}" target="_blank" class="btn btn-primary">View from Yelp</a>
                            <a href="index.html" class="btn btn-default">Back to Search</a>
                            </div>
                        </div>
                    </ul>
                </div>
            </div>
        `;
        $('#restaurant').html(output);
    }).catch((err)=> {
        console.log(err)
    });
}