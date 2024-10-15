import { initializeEditor } from "/static/js/editor.js";
import { initializeRawContent } from "/static/js/raw-content.js";
import { initializePaste } from "./paste.js";

window.onload = () => {
    initializeEditor();
    initializeRawContent();
    initializePaste();
};