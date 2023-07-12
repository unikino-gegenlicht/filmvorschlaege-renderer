<div align="center">
<img src="https://gegenlicht.net/wp-content/themes/gegenlicht/img/logo/gegenlicht_logo_gelb_schwarz.svg">
</div>

# Filmvorschlag Renderer
Dieses Repository enthält eine Software, mit der die Filmvorschläge, die
mithilfe der Survey Monkey Vorlage aufgenommen wurden, in eine HTML-Seite 
umzuwandeln, die anschließend als PDF verteilt werden kann.

Des Weiteren enthält das Repository `*.scss` Dateien, mit denen das Styling
angepasst wird. Hierfür wird ein SASS Compiler benötigt, mit dem zunächst die
`*.scss` Dateien in `*.css` Dateien umgewandelt werden. Hierfür wird in diesem
Projekt `sass` benutzt wird. Unter der Voraussetzung, dass auf dem PC, der für
die Generierung der HTML-Seite verwendet, wird `node` installiert ist, kann mit
```shell
npm install -g sass
```
der aktuelle SCSS Compiler von Dart-Sass installiert werden. Anschliedend kann
mit dem Befehl
```shell
sass style.scss style.css
```
die benötigte CSS-Datei erstellt werden. 