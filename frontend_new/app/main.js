import Backend from "./Backend.js";
import Home from "./views/Home.js";
import About from "./views/About.js";
import Film from "./views/Film.js";
import SearchResults from "./views/SearchResults.js";
// @if DEBUG=true
import SuccessBackendMock from './SuccessBackendMock.js'
// @endif

const pathToRegex = path => {
    return new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");
}

const getParams = match => {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);

    return Object.fromEntries(keys.map((key, i) => {
        return [key, values[i]];
    }));
};

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
};

const router = async () => {
    const routes = [
        { path: "#", view: Home },
        { path: "#/about", view: About },
        { path: "#/film/:id", view: Film },
        { path: "#/search/:query", view: SearchResults }
    ];

    // Test each route for potential match
    const potentialMatches = routes.map(route => {
        return {
            route: route,
            result: location.hash.match(pathToRegex(route.path))
        };
    });

    let match = potentialMatches.find(potentialMatch => potentialMatch.result !== null);

    if (!match) {
        match = {
            route: routes[0],
            result: [location.pathname]
        };
    }

    let view = null;
    // @if DEBUG=true
    view = new match.route.view(new SuccessBackendMock(), getParams(match), document);
    // @endif
    // @if PRODUCTION=true
    view = new match.route.view(new Backend(window.location.hostname), getParams(match), document);
    // @endif 

    document.querySelector("#app").innerHTML = await view.getHtml();
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });

    document.getElementById("global-search-btn").onclick = e => {
        e.preventDefault();
        const query = document.getElementById("global-search-query").value;
        if (query != "" && query != null && query != undefined) {
            navigateTo("/#/search/" + query);
        }
    }

    router();
});
