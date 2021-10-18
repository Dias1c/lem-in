class Graphs {
    static PatternNodeName = `[.\\dA-KM-Za-zА-Яа-я]{1}[.\\d\\wА-Яа-я]{0,}`
    static Palette = { Default: "#1995AD", Start: "#F52549", Finish: "#6af030", Unusing: "#BCBABE", Using: "#1E434C", }

    constructor(content) {
        this.Reconsctruct(content)
    }
    /* == PUBLIC == */
    // Reset All Values By Content
    Reconsctruct(content) {
        this.ClearValues();
        if (!content) {
            return
            // throw { message: "Graph.Reconstructor 'content' undefined or null" }
        }
        let lines = content.split(/\r?\n/)
        // console.log(`Graph Constructor:\nContent = "${content}"\nLines = [${lines}]`)
        lines = this.getNodesWithoutCountAnts(lines)
        // console.log(`After getNodesWithoutCountAnts: ${lines}`)
        lines = this.setNodes(lines)
        lines = this.setLinks(lines)
        if (this.Start != "") {
            this.Steps.push([this.Start])
        }
        lines = this.setSteps(lines)
        this.SetColors()
    }
    // Clear
    ClearValues() {
        this.Start = ""
        this.Finish = ""
        this.Nodes = {}
        this.NodeNames = []
        this.Links = []
        this.Steps = []
        this.Paths = []
    }

    SetColors() {
        // Set Default Colors
        this.NodeNames.forEach((name) => {
            this.Nodes[name].Color = {
                Default: Graphs.Palette.Unusing, Active: Graphs.Palette.Unusing
            }
        });
        let colors = [
            [229, 43, 80],
            [255, 191, 0],
            [0, 127, 255],
            [153, 102, 204],
            [55, 155, 0],
            [0, 49,	83],
            [0, 168, 107],
            [48, 41, 88],
            [245, 120, 0],
            [24, 167, 181],
        ]
        let t = 0
        //Set Colors for Actice
        this.Paths.forEach((path, index) => {
            let r = 0, g = 0, b = 0;
            if(index < colors.length) {
                r = colors[index][0]
                g = colors[index][1]
                b = colors[index][2]
            } else {
                r = getRandom(255);
                g = getRandom(255);
                b = getRandom(255);
            }
            let colorDefault = getRgba(r, g, b, 0.5),
                colorActive = getRgba(r, g, b, 1);
            path.forEach((name) => {
                this.Nodes[name].Color = { Default: colorDefault, Active: colorActive }
            })
        })

        if (this.NodeNames.length != 0) {
            this.Nodes[this.Start].Color = { Default: Graphs.Palette.Start, Active: Graphs.Palette.Start }
            this.Nodes[this.Finish].Color = { Default: Graphs.Palette.Finish, Active: Graphs.Palette.Finish }
        }
    }

    /* == PRIVATE == */
    getNodesWithoutCountAnts(lines) {
        if (lines == null) {
            throw { message: "lines is null, (setNodes)" }
        }
        // Removed line size checker
        let startIdx = 0;
        for (let i = 0; i < lines.length; i++) {
            let line = lines[i]
            startIdx = i + 1
            if (line.startsWith("#") || line == "") { // Check for comment
                if (line.startsWith("##")) {
                    return lines.slice(i)
                }
                continue
            } else {
                let countAnts = parseInt(line);
                if (Number.isInteger(countAnts)) {
                    this.CountAnts = countAnts
                    break
                } else {
                    startIdx--
                    break
                }
            }
        }
        return lines.slice(startIdx)
    }
    setNodes(lines) {
        if (lines == null) {
            throw { message: "lines is null, (setNodes)" }
        }
        // Removed line size checker
        let startIdx = 0;

        // Gets All Rooms
        for (let i = 0; i < lines.length; i++) {
            let line = lines[i]
            startIdx = i + 1
            if (line.startsWith("#")) { // Check for comment or Start|End Rooms
                let isStart = false
                if (line == "##start") {
                    isStart = true
                }
                // check for start or end room
                if (line == "##start" || line == "##end") {
                    i++
                    startIdx++
                    if (i < lines.length) {
                        line = lines[i]
                        let node = this.addNode(line)
                        if (node == null) {
                            throw { message: "after ##start or ##end should be node" }
                        }
                        //Get Room
                        if (node == null) {
                            throw { message: "invalid rooms for start or end not found" }
                        } else if (isStart && this.Start == "") {
                            this.Start = node.Name
                        } else if (!isStart && this.Finish == "") {
                            this.Finish = node.Name
                        } else {
                            throw { message: "there can be only 1 starting and 1 ending rooms on the anthill" }
                        }
                    } else {
                        throw { message: "invalid rooms for start or end not found" }
                    }
                } else {
                    continue
                }
            } else if (line == "") {
                continue
            } else {
                // Take rooms or break if is not valid room
                let node = this.addNode(line)
                if (node == null) {
                    startIdx = startIdx - 1
                    break
                }
            }
        }
        if (this.Start == "" || this.Finish == "") {
            throw { message: "invalid rooms for start or end not found." }
        }
        return lines.slice(startIdx)
    }
    setLinks(lines) {
        if (lines == null) {
            throw { message: "lines is null, (setLinks)" }
        }
        // Removed line size checker
        let startIdx = 0;
        // Gets All Rooms
        for (let i = 0; i < lines.length; i++) {
            let line = lines[i]
            startIdx = i + 1
            if (line.startsWith("#") || line == "") { // Check for comment 
                continue
            } else {
                let link = this.addLink(line)
                if (link == null) {
                    startIdx--
                    break
                } else if (!this.Nodes[link.Source] || !this.Nodes[link.Target]) {
                    throw { message: `Undifined room Name: "${link.Source}" or "${link.Target}"` }
                }
            }
        }
        return lines.slice(startIdx)
    }
    setSteps(lines) {
        if (lines == null) {
            throw { message: "lines is null, (setSteps)" }
        }
        let steps = []
        let startIdx = 0;
        // Gets All Steps
        for (let i = 0; i < lines.length; i++) {
            let line = lines[i]
            startIdx = i + 1
            if (line.startsWith("#") || line == "") { // Check for comment 
                continue
            } else {
                let step = this.addStep(line)
                if (step == null || step.length == 0) {
                    startIdx--
                    break
                }
                steps.push(step)
            }
        }
        // If Steps Not Found Remove Start Room From Step
        if (steps.length == 0) {
            this.Steps = []
            return
        }
        let paths = []
        let ants = {}
        for (let i = 0; i < steps.length; i++) {
            if (i == 0) {
                steps[i].forEach((element, index) => {
                    ants[element.Ant] = { index: index }
                    paths.push([])
                })
            }
            let isBreak = false
            steps[i].forEach((element) => {
                if (ants[element.Ant]) {
                    paths[ants[element.Ant].index].push(element.NodeName)
                }
            })
            if (isBreak) {
                break
            }
        }
        this.Paths = paths
        return lines.slice(startIdx)
    }

    // For Add Node, Gets Node {Name, X, Y}
    addNode(line) {
        if (this.Nodes == null) {
            this.Nodes = {}
        }
        let regex = new RegExp(`^(${Graphs.PatternNodeName}) (-?\\d{1,}) (-?\\d{1,})$`, 'gm')
        let match, node = {};
        if (match = regex.exec(line)) {
            node.Name = match[1],
                node.X = parseInt(match[2]),
                node.Y = parseInt(match[3]);
        } else {
            return null
        }
        if (node != null && node.Name == "") {
            throw { message: "incorrect node: bad name", line }
        }
        this.Nodes[node.Name] = node
        this.NodeNames.push(node.Name)
        return node
    }
    // For Add Link, Gets Link {Sourse, Target}
    addLink(line) {
        if (this.Links == null) {
            this.Links = []
        }
        let regex = new RegExp(`^(${Graphs.PatternNodeName})-(${Graphs.PatternNodeName})$`, 'gm')
        let match, result = {};
        if (match = regex.exec(line)) {
            result.Source = match[1];
            result.Target = match[2];
        } else {
            return null
        }
        this.Links.push(result)
        return result
    }
    // For Add Step, Gets Current Ant Step info {Ant, Node}
    addStep(line) {
        if (this.Steps == null) {
            this.Steps = null
        }
        let regex = new RegExp(`L(\\d{1,})-(${Graphs.PatternNodeName})`, 'gm')
        const match = [...line.matchAll(regex)];
        if (match == null || match.length == 0) {
            return null
        }
        let roomNames = []
        let result = []
        match.forEach(element => {
            roomNames.push(element[2])

            result.push({ Ant: element[1], NodeName: element[2] })
        });
        this.Steps.push(roomNames)
        return result
    }
}

class GraphDrawer {
    Palette = { Default: "#1995AD", Start: "#F52549", Finish: "#6af030", Unusing: "#BCBABE", Using: "#1E434C", } // For Configure
    TileSize = 20
    CircleRadius = 10
    LineWidth = 4
    Transition = 500
    TransitionOnShow = 100
    GraphCircles = {}

    // Graph
    Graph = {}//new Graphs() // {}
    // D3 Element
    html_svg = {}
    html_g = {}
    // 
    Circles = {}

    constructor(blockId, graph) {
        this.Graph = graph
        // this.ConfigureGraph() // if colors value to graph not setted
        this.InitHtmlElements(blockId)
    }

    /* == PUBLIC == */
    InitHtmlElements(blockId) {
        // Clear svg Element in Block
        let html_block = d3.select(blockId)
        html_block.select("svg").remove()
        // Create Elements (INIT)
        this.html_svg = html_block
            .append("svg")
            .attr("viewBox", [0, 0, 500, 500])
            .attr("cursor", "grab");;
        let eventZoom = ({ transform }) => {
            this.html_g.attr("transform", transform);
        }
        this.html_svg
            .call(d3.zoom()
                .on("zoom", eventZoom
                ));
        this.html_g = this.html_svg.append("g")
    }
    ConfigureGraph() {
        // Set Default Colors
        this.Graph.NodeNames.forEach((name) => {
            this.Graph.Nodes[name].Color = {
                Default: this.Palette.Unusing, Active: this.Palette.Unusing
            }
        });
        //Set Colors for Actice
        this.Graph.Paths.forEach((path) => {
            let r = getRandom(255),
                g = getRandom(255),
                b = getRandom(255);
            let colorDefault = getRgba(r, g, b, 0.5),
                colorActive = getRgba(r, g, b, 1);
            path.forEach((name) => {
                this.Graph.Nodes[name].Color = { Default: colorDefault, Active: colorActive }
            })
        })
        if (this.Graph.NodeNames.length != 0) {
            this.Graph.Nodes[this.Graph.Start].Color = { Default: this.Palette.Start, Active: this.Palette.Start }
            this.Graph.Nodes[this.Graph.Finish].Color = { Default: this.Palette.Finish, Active: this.Palette.Finish }
        }
    }
    ReDrawGraph() {
        this.html_g.html("");
        this.AddGraphToGroup();
    }
    AddGraphToGroup() {
        this.addLines();
        this.addMovingLines();
        this.addCircles();
        this.addCirlceNames();
    }
    ShowFrame(frame) {
        if (0 <= frame && frame < this.Graph.Steps.length) {

            let usedNodes = {}
            this.Graph.Steps[frame].forEach((name) => {
                this.Circles[name]
                    .transition()
                    .duration(this.Transition)
                    .attr("r", this.CircleRadius + (this.CircleRadius * 0.2))
                    .style("fill", this.Graph.Nodes[name].Color.Active)
                usedNodes[name] = true
            })

            this.Graph.NodeNames.forEach((name) => {
                if (!usedNodes[name]) {
                    this.Circles[name].transition()
                        .duration(this.Transition)
                        .attr("r", this.CircleRadius)
                        .style("fill", this.Graph.Nodes[name].Color.Default)
                }
            })
        }
    }
    /* == PRIVATE == */
    addLines() {
        this.Graph.Links.forEach((node) => {
            let from = this.Graph.Nodes[node.Source]
            let to = this.Graph.Nodes[node.Target]

            this.html_g
                .append("line")
                .style("stroke-linecap", "round")
                .style("stroke-width", this.LineWidth)
                .style("stroke", "#dfdfdf")
                .attr("x1", from.X * this.TileSize).attr("y1", from.Y * this.TileSize)
                .attr("x2", to.X * this.TileSize).attr("y2", to.Y * this.TileSize);
        })
    }
    addMovingLines() {
        this.Graph.Paths.forEach((path)=>{
            let prev = this.Graph.Nodes[this.Graph.Start]
            let lineWidth = this.LineWidth/2

            let frame = 1
            path.forEach((name) => {
                let node = this.Graph.Nodes[name]
                let stroke = this.html_g
                    .append("line")
                    .attr("x1", prev.X * this.TileSize).attr("y1", prev.Y * this.TileSize)
                    .attr("x2", node.X * this.TileSize).attr("y2", node.Y * this.TileSize)
                    .style("stroke-linecap", "round")
                    .style("stroke-width", lineWidth)
                    .style("stroke", "#dfdfdf")
                    .transition()
                    .duration(this.TransitionOnShow*frame)
                    .style("stroke", node.Color.Default);
                if (node.Name == this.Graph.Finish) {
                    stroke.style("stroke", prev.Color.Default);
                }
                prev = node
                frame++
            });
        });
    }
    addCircles() {
        // let node = "nodeName"
        this.Graph.NodeNames.forEach((name) => {
            let node = this.Graph.Nodes[name]
            let posX = node.X * this.TileSize
            let posY = node.Y * this.TileSize

            this.Circles[name] = this.html_g
                .append("circle")
                .attr("cx", posX)
                .attr("cy", posY)
                .attr("r", this.CircleRadius)
                .style("fill", node.Color.Default)
        });
        if (this.Circles[this.Graph.Start]) {
            this.Circles[this.Graph.Start].style("stroke", "#000")
        }
        if (this.Circles[this.Graph.Finish]) {
            this.Circles[this.Graph.Finish].style("stroke", "#000")
        }
    }
    addCirlceNames() {
        let fontSize = this.CircleRadius*1.5
        this.Graph.NodeNames.forEach((name) => {
            let node = this.Graph.Nodes[name]
            let posX = node.X * this.TileSize
            let posY = node.Y * this.TileSize

            this.html_g
                .append("text")
                .attr("x", posX)
                .attr("y", posY - (this.CircleRadius + 5))
                .attr("fill", "#000")
                .style("font-size", `${fontSize}px`)
                .text(`${node.Name}`)
                .attr("text-anchor", "middle")
        });
    }
}



// Utils
function getRandom(max) {
    let o = Math.round, r = Math.random;
    return o(r() * max)
}

function getRgba(r, g, b, opacity) {
    return `rgba(${r}, ${g}, ${b}, ${opacity})`;
}