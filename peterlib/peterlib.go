package peterlib

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	Cmd   string
	Value interface{}
}

var log []Command

const canvasWidth = 400
const canvasHeight = 400

// -------------------------------------------------
// Fonctions de contrôle de la tortue
// -------------------------------------------------

func Down() {
	fmt.Println("Action: Down (stylo baissé)")
	log = append(log, Command{"Down", nil})
}

func Up() {
	fmt.Println("Action: Up (stylo levé)")
	log = append(log, Command{"Up", nil})
}

func Right() {
	fmt.Println("Action: Right 90°")
	log = append(log, Command{"Right", 90.0})
}

func Pivote(angle float64) {
	fmt.Printf("Action: Right %.2f°\n", angle)
	log = append(log, Command{"Right", angle})
}

func Left() {
	fmt.Println("Action: Left 90°")
	log = append(log, Command{"Left", 90.0})
}

func SetHeading(angle float64) {
	fmt.Printf("Action: SetHeading à %.2f°\n", angle)
	log = append(log, Command{"SetHeading", angle})
}

// Fonctions d'orientation absolue
func North() {
	fmt.Println("Action: Orientation Nord (0°)")
	SetHeading(0)
}

func East() {
	fmt.Println("Action: Orientation Est (90°)")
	SetHeading(90)
}

func South() {
	fmt.Println("Action: Orientation Sud (180°)")
	SetHeading(180)
}

func West() {
	fmt.Println("Action: Orientation Ouest (270°)")
	SetHeading(270)
}

func Forward(n int) {
	fmt.Printf("Action: Forward %d px\n", n)
	log = append(log, Command{"Forward", float64(n)})
}

func Backward(n float64) {
	fmt.Printf("Action: Backward %.2f px\n", n)
	log = append(log, Command{"Backward", n})
}

func GoTo(x, y float64) {
	fmt.Printf("Action: GoTo (%.2f, %.2f)\n", x, y)
	log = append(log, Command{"GoTo", []float64{x, y}})
}

func Color(c string) {
	fmt.Printf("Action: Color %s\n", c)
	log = append(log, Command{"Color", c})
}

func PenSize(width float64) {
	fmt.Printf("Action: PenSize %.2f\n", width)
	log = append(log, Command{"PenSize", width})
}

func Circle(radius float64) {
	fmt.Printf("Action: Circle rayon %.2f\n", radius)
	log = append(log, Command{"Circle", radius})
}

func Say(msg string) {
	fmt.Printf("Action: Say \"%s\"\n", msg)
	log = append(log, Command{"Say", msg})
}

// Play génère le HTML et lance l’animation
func Play() {
	if len(log) > 0 {
		const filename = "peter.html"
		generateHTML(filename)
		openBrowser(filename)
	} else {
		fmt.Println("Aucune commande à exécuter.")
	}
}

// openBrowser ouvre l'URL spécifiée dans le navigateur par défaut.
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Run()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Run()
	case "darwin":
		err = exec.Command("open", url).Run()
	default:
		err = fmt.Errorf("plateforme non supportée pour l'ouverture automatique du navigateur")
	}

	if err != nil {
		fmt.Printf("\nImpossible d'ouvrir le navigateur automatiquement: %v\n", err)
		fmt.Printf("Veuillez ouvrir le fichier '%s' manuellement.\n", url)
	}
}

// -------------------------------------------------
// Génération HTML
// -------------------------------------------------
func generateHTML(filename string) {
	// Convertir les commandes en une chaîne JSON
	var cmdsJs strings.Builder
	cmdsJs.WriteString("[\n")
	for i, c := range log {
		// Utilisation de %q pour les chaînes de caractères pour gérer correctement les guillemets
		if val, ok := c.Value.(string); ok {
			fmt.Fprintf(&cmdsJs, `    {cmd:"%s", value:%q}`, c.Cmd, val)
		} else if c.Value != nil {
			fmt.Fprintf(&cmdsJs, `    {cmd:"%s", value:%v}`, c.Cmd, c.Value)
		} else {
			fmt.Fprintf(&cmdsJs, `    {cmd:"%s"}`, c.Cmd)
		}
		if i < len(log)-1 {
			cmdsJs.WriteString(",\n")
		}
	}
	cmdsJs.WriteString("\n]")

	// Utilisation d'un template pour la génération HTML
	t, err := template.New("peter").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Erreur parsing template:", err)
		return
	}

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Erreur création fichier %s: %v\n", filename, err)
		return
	}
	defer f.Close()

	data := struct {
		CanvasWidth  int
		CanvasHeight int
		Commands     template.JS
	}{
		CanvasWidth:  canvasWidth,
		CanvasHeight: canvasHeight,
		Commands:     template.JS(cmdsJs.String()),
	}

	err = t.Execute(f, data)
	if err != nil {
		fmt.Println("Erreur exécution template:", err)
	}
}

const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Peter le Traceur</title>
<style>
  html, body {
    margin: 0;
    padding: 0;
    height: 100%;
    overflow: hidden; /* Empêche les barres de défilement */
    background-color: #333;
  }
  canvas {
    --grid-size: 37.8px; /* Approx. 1cm at 96 DPI */
    display: block;
    background-color:#f9f9f9;
    background-image:
      linear-gradient(to right, #ddd 1px, transparent 1px),
      linear-gradient(to bottom, #ddd 1px, transparent 1px);
    background-size: var(--grid-size) var(--grid-size);
  }
</style>
</head>
<body>
<canvas id="c"></canvas>
<script>
const cmds = {{.Commands}};
const canvas = document.getElementById("c");
const cmToPx = 37.795; // Conversion factor: 1cm to pixels (at 96 DPI)
const ctx = canvas.getContext("2d");

// Ajuster la taille du canevas à la fenêtre
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

// Position initiale au centre du canevas
const gridSize = parseFloat(getComputedStyle(canvas).getPropertyValue('--grid-size'));
let x = Math.round(canvas.width / 2 / gridSize) * gridSize + 0.5;
let y = Math.round(canvas.height / 2 / gridSize) * gridSize + 0.5;

let angle = -90; // -90 = vers le haut
let penDown = false;
ctx.lineWidth = 2;
//ctx.strokeStyle = "black";

function drawTurtle() {
    // Dessine un bonhomme stylisé
    ctx.save();
    ctx.translate(x, y);
    ctx.rotate(angle * Math.PI / 180);
    // ctx.strokeStyle = "black"; // SI : pour pouvoir changer la couleur
    ctx.lineWidth = 2;

    ctx.beginPath();
    // Tête
    ctx.arc(0, -12, 4, 0, 2 * Math.PI);
    // Corps
    ctx.moveTo(0, -8);
    ctx.lineTo(0, 4);
    // Bras
    ctx.moveTo(-7, 0);
    ctx.lineTo(7, 0);
    // Jambes
    ctx.moveTo(0, 4);
    ctx.lineTo(-5, 12);
    ctx.moveTo(0, 4);
    ctx.lineTo(5, 12);
    ctx.stroke();
    ctx.restore();
}

let i = 0;
let intervalId = null;
let speed = 100;
let path = []; // Pour stocker les segments du chemin

let couleur = "black"

function step(){
    if(i >= cmds.length){
        clearInterval(intervalId);
        intervalId = null;
        redrawAll(); // Un dernier dessin pour s'assurer que tout est à jour
        return;
    }
    let c = cmds[i++];
    
    const oldX = x;
    const oldY = y;

    switch(c.cmd){
        case "Down": penDown = true; break;
        case "Up": penDown = false; break;
        case "Right": angle += parseFloat(c.value); break;
        case "Left": angle -= 90; break;
        case "SetHeading": angle = parseFloat(c.value) - 90; break; // Ajustement pour que 0 soit vers le haut
        case "Forward":
            x += c.value * cmToPx * Math.cos(angle * Math.PI / 180);
            y += c.value * cmToPx * Math.sin(angle * Math.PI / 180);
            break;
        case "Backward":
            x -= c.value * cmToPx * Math.cos(angle * Math.PI / 180);
            y -= c.value * cmToPx * Math.sin(angle * Math.PI / 180);
            break;
        case "GoTo":
            x = (canvas.width / 2) + c.value[0] * cmToPx;
            y = (canvas.height / 2) - c.value[1] * cmToPx;
            break;
        case "Color": couleur = c.value; break;
        case "PenSize": ctx.lineWidth = parseFloat(c.value); break;
        case "Circle":
            const radiusPx = c.value * cmToPx;
            // Pour les cercles, on ajoute un segment spécial au chemin
            if (penDown) {
                path.push({type: 'circle', x: x, y: y, radius: radiusPx, angle: angle, style: ctx.strokeStyle, width: ctx.lineWidth});
            }
            break;
        case "Say": setTimeout(()=>alert("Peter dit: "+c.value),50); break;
    }

    // Si le stylo était baissé et qu'un mouvement a eu lieu, on ajoute un segment au chemin
    if (penDown && (c.cmd === "Forward" || c.cmd === "Backward" || c.cmd === "GoTo")) {
        path.push({type: 'line', x1: oldX, y1: oldY, x2: x, y2: y, style: couleur, width: ctx.lineWidth});
    }

    redrawAll();
}

function redrawAll() {
    // Efface le canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Redessine le chemin parcouru
    for (const segment of path) {
        ctx.strokeStyle = segment.style;
        ctx.lineWidth = segment.width;
        ctx.beginPath();
        if (segment.type === 'line') {
            ctx.moveTo(segment.x1, segment.y1);
            ctx.lineTo(segment.x2, segment.y2);
        } else if (segment.type === 'circle') {
            ctx.arc(segment.x + segment.radius * Math.sin(segment.angle * Math.PI / 180), segment.y - segment.radius * Math.cos(segment.angle * Math.PI / 180), segment.radius, 0, 2 * Math.PI);
        }
        ctx.stroke();
    }

    // Dessine le bonhomme à sa position actuelle
    drawTurtle();
}

function play(){
    if(intervalId) clearInterval(intervalId);
    intervalId = setInterval(step, speed);
}

// Optionnel : Gérer le redimensionnement de la fenêtre (recharge la page pour redessiner)
window.onresize = function() {
    location.reload();
}

// Démarrage automatique
play();

</script>
</body>
</html>
`
