// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file visualization/ssl_visualization.proto (syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file visualization/ssl_visualization.proto.
 */
export const file_visualization_ssl_visualization: GenFile = /*@__PURE__*/
  fileDesc("CiV2aXN1YWxpemF0aW9uL3NzbF92aXN1YWxpemF0aW9uLnByb3RvIjYKCFJnYkNvbG9yEgkKAXIYASABKA0SCQoBZxgCIAEoDRIJCgFiGAMgASgNEgkKAWEYBCABKAIiggEKCE1ldGFkYXRhEg0KBWxheWVyGAEgAygJEhgKEHZpc2libGVCeURlZmF1bHQYAiABKAgSDQoFb3JkZXIYAyABKAUSHQoKY29sb3JfZmlsbBgEIAEoCzIJLlJnYkNvbG9yEh8KDGNvbG9yX3N0cm9rZRgFIAEoCzIJLlJnYkNvbG9yImoKC0xpbmVTZWdtZW50EhsKCG1ldGFkYXRhGAEgASgLMgkuTWV0YWRhdGESDwoHc3RhcnRfeBgCIAEoAhIPCgdzdGFydF95GAMgASgCEg0KBWVuZF94GAQgASgCEg0KBWVuZF95GAUgASgCIlkKBkNpcmNsZRIbCghtZXRhZGF0YRgBIAEoCzIJLk1ldGFkYXRhEhAKCGNlbnRlcl94GAIgASgCEhAKCGNlbnRlcl95GAMgASgCEg4KBnJhZGl1cxgEIAEoAiJeChJWaXN1YWxpemF0aW9uRnJhbWUSEQoJc2VuZGVyX2lkGAEgASgJEhsKBWxpbmVzGAIgAygLMgwuTGluZVNlZ21lbnQSGAoHY2lyY2xlcxgDIAMoCzIHLkNpcmNsZUJaQhVTc2xWaXN1YWxpemF0aW9uUHJvdG9QAVo/Z2l0aHViLmNvbS9Sb2JvQ3VwLVNTTC9zc2wtdmlzaW9uLWNsaWVudC9pbnRlcm5hbC92aXN1YWxpemF0aW9uYgZwcm90bzM");

/**
 * Color encoded in RGB
 *
 * @generated from message RgbColor
 */
export type RgbColor = Message<"RgbColor"> & {
  /**
   * red (0-255)
   *
   * @generated from field: uint32 r = 1;
   */
  r: number;

  /**
   * green (0-255)
   *
   * @generated from field: uint32 g = 2;
   */
  g: number;

  /**
   * blue (0-255)
   *
   * @generated from field: uint32 b = 3;
   */
  b: number;

  /**
   * alpha (0.0-1.0)
   *
   * @generated from field: float a = 4;
   */
  a: number;
};

/**
 * Color encoded in RGB
 *
 * @generated from message RgbColor
 */
export type RgbColorJson = {
  /**
   * red (0-255)
   *
   * @generated from field: uint32 r = 1;
   */
  r?: number;

  /**
   * green (0-255)
   *
   * @generated from field: uint32 g = 2;
   */
  g?: number;

  /**
   * blue (0-255)
   *
   * @generated from field: uint32 b = 3;
   */
  b?: number;

  /**
   * alpha (0.0-1.0)
   *
   * @generated from field: float a = 4;
   */
  a?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message RgbColor.
 * Use `create(RgbColorSchema)` to create a new message.
 */
export const RgbColorSchema: GenMessage<RgbColor, RgbColorJson> = /*@__PURE__*/
  messageDesc(file_visualization_ssl_visualization, 0);

/**
 * Metadata for each shape
 *
 * @generated from message Metadata
 */
export type Metadata = Message<"Metadata"> & {
  /**
   * layer name, optionally with a hierarchy
   *
   * @generated from field: repeated string layer = 1;
   */
  layer: string[];

  /**
   * Should a client show this by default?
   *
   * @generated from field: bool visibleByDefault = 2;
   */
  visibleByDefault: boolean;

  /**
   * An order number:
   * <0: Below field lines
   * 0: default
   * 1: robots
   * 2: robot ids
   * 3: ball
   * >3: above vision objects
   *
   * @generated from field: int32 order = 3;
   */
  order: number;

  /**
   * Color to fill the shape
   *
   * @generated from field: RgbColor color_fill = 4;
   */
  colorFill?: RgbColor;

  /**
   * Color for the shape stroke
   *
   * @generated from field: RgbColor color_stroke = 5;
   */
  colorStroke?: RgbColor;
};

/**
 * Metadata for each shape
 *
 * @generated from message Metadata
 */
export type MetadataJson = {
  /**
   * layer name, optionally with a hierarchy
   *
   * @generated from field: repeated string layer = 1;
   */
  layer?: string[];

  /**
   * Should a client show this by default?
   *
   * @generated from field: bool visibleByDefault = 2;
   */
  visibleByDefault?: boolean;

  /**
   * An order number:
   * <0: Below field lines
   * 0: default
   * 1: robots
   * 2: robot ids
   * 3: ball
   * >3: above vision objects
   *
   * @generated from field: int32 order = 3;
   */
  order?: number;

  /**
   * Color to fill the shape
   *
   * @generated from field: RgbColor color_fill = 4;
   */
  colorFill?: RgbColorJson;

  /**
   * Color for the shape stroke
   *
   * @generated from field: RgbColor color_stroke = 5;
   */
  colorStroke?: RgbColorJson;
};

/**
 * Describes the message Metadata.
 * Use `create(MetadataSchema)` to create a new message.
 */
export const MetadataSchema: GenMessage<Metadata, MetadataJson> = /*@__PURE__*/
  messageDesc(file_visualization_ssl_visualization, 1);

/**
 * A line segment
 *
 * @generated from message LineSegment
 */
export type LineSegment = Message<"LineSegment"> & {
  /**
   * The metadata
   *
   * @generated from field: Metadata metadata = 1;
   */
  metadata?: Metadata;

  /**
   * Start point, x value [m]
   *
   * @generated from field: float start_x = 2;
   */
  startX: number;

  /**
   * Start point, y value [m]
   *
   * @generated from field: float start_y = 3;
   */
  startY: number;

  /**
   * End point, x value [m]
   *
   * @generated from field: float end_x = 4;
   */
  endX: number;

  /**
   * End point, y value [m]
   *
   * @generated from field: float end_y = 5;
   */
  endY: number;
};

/**
 * A line segment
 *
 * @generated from message LineSegment
 */
export type LineSegmentJson = {
  /**
   * The metadata
   *
   * @generated from field: Metadata metadata = 1;
   */
  metadata?: MetadataJson;

  /**
   * Start point, x value [m]
   *
   * @generated from field: float start_x = 2;
   */
  startX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Start point, y value [m]
   *
   * @generated from field: float start_y = 3;
   */
  startY?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * End point, x value [m]
   *
   * @generated from field: float end_x = 4;
   */
  endX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * End point, y value [m]
   *
   * @generated from field: float end_y = 5;
   */
  endY?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message LineSegment.
 * Use `create(LineSegmentSchema)` to create a new message.
 */
export const LineSegmentSchema: GenMessage<LineSegment, LineSegmentJson> = /*@__PURE__*/
  messageDesc(file_visualization_ssl_visualization, 2);

/**
 * A full circle
 *
 * @generated from message Circle
 */
export type Circle = Message<"Circle"> & {
  /**
   * The metadata
   *
   * @generated from field: Metadata metadata = 1;
   */
  metadata?: Metadata;

  /**
   * Center point, x value [m]
   *
   * @generated from field: float center_x = 2;
   */
  centerX: number;

  /**
   * Center point, y value [m]
   *
   * @generated from field: float center_y = 3;
   */
  centerY: number;

  /**
   * Radius [m]
   *
   * @generated from field: float radius = 4;
   */
  radius: number;
};

/**
 * A full circle
 *
 * @generated from message Circle
 */
export type CircleJson = {
  /**
   * The metadata
   *
   * @generated from field: Metadata metadata = 1;
   */
  metadata?: MetadataJson;

  /**
   * Center point, x value [m]
   *
   * @generated from field: float center_x = 2;
   */
  centerX?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Center point, y value [m]
   *
   * @generated from field: float center_y = 3;
   */
  centerY?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * Radius [m]
   *
   * @generated from field: float radius = 4;
   */
  radius?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message Circle.
 * Use `create(CircleSchema)` to create a new message.
 */
export const CircleSchema: GenMessage<Circle, CircleJson> = /*@__PURE__*/
  messageDesc(file_visualization_ssl_visualization, 3);

/**
 * Wrapper frame containing all shapes
 *
 * @generated from message VisualizationFrame
 */
export type VisualizationFrame = Message<"VisualizationFrame"> & {
  /**
   * An identifier for the sender
   * Used to identify the source of shapes in a client
   * Also used to keep track of the latest frame of each sender in clients, if there a multiple ones senders
   *
   * @generated from field: string sender_id = 1;
   */
  senderId: string;

  /**
   * all lines
   *
   * @generated from field: repeated LineSegment lines = 2;
   */
  lines: LineSegment[];

  /**
   * all circles
   *
   * @generated from field: repeated Circle circles = 3;
   */
  circles: Circle[];
};

/**
 * Wrapper frame containing all shapes
 *
 * @generated from message VisualizationFrame
 */
export type VisualizationFrameJson = {
  /**
   * An identifier for the sender
   * Used to identify the source of shapes in a client
   * Also used to keep track of the latest frame of each sender in clients, if there a multiple ones senders
   *
   * @generated from field: string sender_id = 1;
   */
  senderId?: string;

  /**
   * all lines
   *
   * @generated from field: repeated LineSegment lines = 2;
   */
  lines?: LineSegmentJson[];

  /**
   * all circles
   *
   * @generated from field: repeated Circle circles = 3;
   */
  circles?: CircleJson[];
};

/**
 * Describes the message VisualizationFrame.
 * Use `create(VisualizationFrameSchema)` to create a new message.
 */
export const VisualizationFrameSchema: GenMessage<VisualizationFrame, VisualizationFrameJson> = /*@__PURE__*/
  messageDesc(file_visualization_ssl_visualization, 4);

