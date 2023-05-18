document.getElementById("check-availability-button").addEventListener("click", ()=>{
    let html = `
        <form id="check-availability-button" action="" method="post" class="overflow-hidden needs-validation" novalidate>
            <div class="row">
                <div class="col">
                    <div class="row" id="reservation-dates-modal">
                        <div class="col">
                            <input required disabled class="form-control" type="text" name="start" id="start" placeholder="Arrival">
                        </div>
                        <div class="col">
                            <input required disabled class="form-control" type="text" name="end" id="end" placeholder="Departure">
                        </div>
                    </div>

                </div>
            </div>
        </form>
        `;
        attention.custom({msg: html, title: "Choose your date"})
})