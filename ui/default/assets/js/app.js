import { initializeEditor } from "/static/js/editor.js";
import { initializePaste } from "./paste.js";

window.onload = () => {
    initializeEditor();
    initializePaste();
};