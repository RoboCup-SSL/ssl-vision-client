const fieldWidth = 9000;
const fieldLength = 12000;
const centerCircleRadius = 500;

export interface Point {
    x: number,
    y: number,
}

export interface LineShape {
    p1: Point,
    p2: Point,
}

export interface CircleShape {
    center: Point,
    radius: number,
    stroke: string,
    strokeWidth: number,
    fill: string,
    fillOpacity: number,
}

export interface TextShape {
    text: string,
    p: Point,
    fill: string,
}

export interface PathSegment {
    type: 'M' | 'A' | 'L'
    args: number[],
}

export interface PathShape {
    d: PathSegment[],
    stroke: string,
    strokeWidth: number,
    fill: string,
    fillOpacity: number
}

export type Shape = { line: LineShape } | { circle: CircleShape } | { text: TextShape } | { path: PathShape }

export interface Field {
    activeSourceId: string,
    sources: any,
    fieldWidth: number,
    fieldLength: number,
    boundaryWidth: number,
    penAreaWidth: number,
    penAreaDepth: number,
    goalWidth: number,
    goalDepth: number,
    centerCircleRadius: number,
    ballRadius: number,
    shapes: Shape[],
}

export const defaultField: Field = {
    activeSourceId: '',
    sources: {},
    fieldWidth: fieldWidth,
    fieldLength: fieldLength,
    boundaryWidth: 300,
    penAreaWidth: 1000,
    penAreaDepth: 500,
    goalWidth: 600,
    goalDepth: 180,
    centerCircleRadius: centerCircleRadius,
    ballRadius: 21.5,
    shapes: [
        {
            line: {
                p1: {x: -fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: fieldWidth / 2}
            },
        },
        {

            line: {
                p1: {x: -fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: -fieldLength / 2, y: fieldWidth / 2},
                p2: {x: fieldLength / 2, y: fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: fieldLength / 2, y: fieldWidth / 2},
                p2: {x: fieldLength / 2, y: -fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: fieldLength / 2, y: -fieldWidth / 2},
                p2: {x: -fieldLength / 2, y: -fieldWidth / 2}
            }
        },
        {
            line: {
                p1: {x: 0, y: fieldWidth / 2},
                p2: {x: 0, y: -fieldWidth / 2}
            }
        },
        {
            circle: {
                center: {x: 0, y: 0},
                radius: centerCircleRadius,
                stroke: 'white',
                strokeWidth: 10,
                fill: '',
                fillOpacity: 0
            }
        },
        {
            path: {
                d: [
                    {
                        type: 'M',
                        args: [925, -550]
                    },
                    {
                        type: 'A',
                        args: [90, 90, 0, 1, 1, 925, -450]
                    },
                    {
                        type: 'L',
                        args: [925, -550]
                    }
                ],
                stroke: 'black',
                strokeWidth: 10,
                fill: 'yellow',
                fillOpacity: 1
            }
        },
        {
            text: {
                text: '1',
                p: {x: 990, y: -510},
                fill: 'black',
            }
        }
    ],
}
