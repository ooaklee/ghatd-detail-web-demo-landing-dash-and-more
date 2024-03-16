<p align="center">
  <a href="https://ghatd.com">
    <img alt="GHATD" src="https://demo.ghatd.com/static/img/ghat-logo-square.png" width="180" />
  </a>
</p>
<h1 align="center">
  GHAT(D)'s Demo <code>Web Detail</code> - Landing, Dashboard, and more [WIP]
</h1>
<h4 align="center">
  Apart of the <a href="http://github.com/ooaklee/ghatd" target="_blank">GHAT(D) initiative</a>
</h4>


> This repo is strongly linked to this [**GHAT(D) PR** being merged](https://github.com/ooaklee/ghatd/pull/2)

Use GHAT(D) `Details` to hit the floor running with your next Go-base full stack web application. This `Detail` comes loaded with everything seen on [**our demo page**](https://demo.ghatd.com). Use this demo to 
Kick off your project if you want to get an understanding of how you can structure the **web `Detail`** of your application


## 🚥 Getting started

`Details` are independent applications by nature that can run within the GHAT(D) framework as well as individually. To run this `web` Detail locally, please:

- Ensure you have the appropiate version of Go installed
- Run the following command:
```sh
go run web.go
```
- By default you should be able to access your `web` Detail on http://localhost:4044


For the best experience developing your Detail, we recommend using hot reloading when developing by using:
```sh
reflex -r '\.(html|go|css|png|svg|ico|js|woff2|woff|ttf|eot)$' -s -- go run web.go
```
You will have to ensure you have [**reflex** (click to go to installation steps)](https://github.com/cespare/reflex?tab=readme-ov-file#installation) installed.

## 🪡 Putting together your web application and the ghat(d) framework (TBC)

1.  **Create a GHAT(D) Web App.**

    Use the GHATDCLI to create a new web app, specifying this demo `web` detail.

    ```shell
    # create a new GHAT(D) Web App using this demo web detail
    ghatdcli new -n "my-new-web-app" -m "github.com/some-org-or-personal/my-new-web-app" -w "https://github.com/ooaklee/ghatd-detail-web-demo-landing-dash-and-more"
    ```

2.  **Start developing.**

    Navigate into your new web app's directory and start it up.

    ```shell
    cd my-new-web-app/
    go mod tidy
    go run cmd/main.go start-server
    ```

    > For the best experience, we recommend using hot reloading when developing by using `reflex -r '\.(html|go|css|png|svg|ico|js|woff2|woff|ttf|eot)$' -s -- go run main.go start-server`. You will have to ensure you have 
    [**reflex** (click to go to installation steps)](https://github.com/cespare/reflex?tab=readme-ov-file#installation) installed.
    

3.  **Browse the service and and start editing your app code!**

    Your site is now running at `http://localhost:4000`!


## More about this `web` Detail

See below for more information on the core components used for this `Detail`'s stack.

- **go:** [v1.22.x](https://go.dev/doc/install)
- **htmx:** [v1.9.10](https://htmx.org/)
- **alpine.js:** [v3.x](https://alpinejs.dev/essentials/installation#from-a-script-tag)
- **tailwindcss:** [v3.x](https://github.com/asdf-community/asdf-golang)
- **daisy ui:** [v3.x](https://daisyui.com/docs/install/)
  - Notable alternatives include:
    - **flowbite:** [v2.3.x](https://flowbite.com/docs/getting-started/introduction/#include-via-cdn)
    - **wind-ui:** [v.3.4.x](https://wind-ui.com/)
- **version manager:** [asdf](https://github.com/asdf-vm/asdf)

> The dashboard's base template was taken from the [**`TailAdmin` team**](https://tailadmin.com/). Please support them by [**purchasing their templates**](https://tailadmin.com/pricing) or giving their [**GitHub repository**](https://github.com/TailAdmin/tailadmin-free-tailwind-dashboard-template) a star.


## License
This project is licensed under the [MIT License](./LICENSE).
