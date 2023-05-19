document.getElementById("check-availability-button").addEventListener("click", ()=>{
    let html = `
        <form id="check-availability-form" class="overflow-hidden needs-validation" novalidate>
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
        attention.custom({
            msg: html, 
            title: "Choose your date",
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    showOnFocus: true,
                })
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled')
                document.getElementById('end').removeAttribute('disabled')
            },
            callback: async function(result) {
                console.log("called");

                let form = document.getElementById("check-availability-form");
                let formData = new FormData(form);
                let tkn = document.querySelector('meta[name="csrf_token"]').content
                formData.append("csrf_token", tkn);
                console.log(formData.get("csrf_token"))

                const resp = await fetch('/search-avaliable-json', {
                    method: "POST",
                    body: formData,
                });
                const data = await resp.json();
                console.log(data)
            }
        })
})