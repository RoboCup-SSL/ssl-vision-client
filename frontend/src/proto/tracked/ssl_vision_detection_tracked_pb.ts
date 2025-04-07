// @generated by protoc-gen-es v2.2.5 with parameter "target=ts,json_types=true"
// @generated from file tracked/ssl_vision_detection_tracked.proto (syntax proto2)
 

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { RobotId, RobotIdJson } from "../gc/ssl_gc_common_pb";
import { file_gc_ssl_gc_common } from "../gc/ssl_gc_common_pb";
import type { Vector2, Vector2Json, Vector3, Vector3Json } from "../gc/ssl_gc_geometry_pb";
import { file_gc_ssl_gc_geometry } from "../gc/ssl_gc_geometry_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file tracked/ssl_vision_detection_tracked.proto.
 */
export const file_tracked_ssl_vision_detection_tracked: GenFile = /*@__PURE__*/
  fileDesc("Cip0cmFja2VkL3NzbF92aXNpb25fZGV0ZWN0aW9uX3RyYWNrZWQucHJvdG8iTwoLVHJhY2tlZEJhbGwSFQoDcG9zGAEgAigLMgguVmVjdG9yMxIVCgN2ZWwYAiABKAsyCC5WZWN0b3IzEhIKCnZpc2liaWxpdHkYAyABKAIiowEKCktpY2tlZEJhbGwSFQoDcG9zGAEgAigLMgguVmVjdG9yMhIVCgN2ZWwYAiACKAsyCC5WZWN0b3IzEhcKD3N0YXJ0X3RpbWVzdGFtcBgDIAIoARIWCg5zdG9wX3RpbWVzdGFtcBgEIAEoARIaCghzdG9wX3BvcxgFIAEoCzIILlZlY3RvcjISGgoIcm9ib3RfaWQYBiABKAsyCC5Sb2JvdElkIpYBCgxUcmFja2VkUm9ib3QSGgoIcm9ib3RfaWQYASACKAsyCC5Sb2JvdElkEhUKA3BvcxgCIAIoCzIILlZlY3RvcjISEwoLb3JpZW50YXRpb24YAyACKAISFQoDdmVsGAQgASgLMgguVmVjdG9yMhITCgt2ZWxfYW5ndWxhchgFIAEoAhISCgp2aXNpYmlsaXR5GAYgASgCIrgBCgxUcmFja2VkRnJhbWUSFAoMZnJhbWVfbnVtYmVyGAEgAigNEhEKCXRpbWVzdGFtcBgCIAIoARIbCgViYWxscxgDIAMoCzIMLlRyYWNrZWRCYWxsEh0KBnJvYm90cxgEIAMoCzINLlRyYWNrZWRSb2JvdBIgCgtraWNrZWRfYmFsbBgFIAEoCzILLktpY2tlZEJhbGwSIQoMY2FwYWJpbGl0aWVzGAYgAygOMgsuQ2FwYWJpbGl0eSqSAQoKQ2FwYWJpbGl0eRIWChJDQVBBQklMSVRZX1VOS05PV04QABIiCh5DQVBBQklMSVRZX0RFVEVDVF9GTFlJTkdfQkFMTFMQARIkCiBDQVBBQklMSVRZX0RFVEVDVF9NVUxUSVBMRV9CQUxMUxACEiIKHkNBUEFCSUxJVFlfREVURUNUX0tJQ0tFRF9CQUxMUxADQl1CHlNzbFZpc2lvbkRldGVjdGlvblRyYWNrZWRQcm90b1ABWjlnaXRodWIuY29tL1JvYm9DdXAtU1NML3NzbC12aXNpb24tY2xpZW50L2ludGVybmFsL3RyYWNrZWQ", [file_gc_ssl_gc_common, file_gc_ssl_gc_geometry]);

/**
 * A single tracked ball
 *
 * @generated from message TrackedBall
 */
export type TrackedBall = Message<"TrackedBall"> & {
  /**
   * The position (x, y, height) [m] in the ssl-vision coordinate system
   *
   * @generated from field: required Vector3 pos = 1;
   */
  pos?: Vector3;

  /**
   * The velocity [m/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional Vector3 vel = 2;
   */
  vel?: Vector3;

  /**
   * The visibility of the ball
   * A value between 0 (not visible) and 1 (visible)
   * The exact implementation depends on the source software
   *
   * @generated from field: optional float visibility = 3;
   */
  visibility: number;
};

/**
 * A single tracked ball
 *
 * @generated from message TrackedBall
 */
export type TrackedBallJson = {
  /**
   * The position (x, y, height) [m] in the ssl-vision coordinate system
   *
   * @generated from field: required Vector3 pos = 1;
   */
  pos?: Vector3Json;

  /**
   * The velocity [m/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional Vector3 vel = 2;
   */
  vel?: Vector3Json;

  /**
   * The visibility of the ball
   * A value between 0 (not visible) and 1 (visible)
   * The exact implementation depends on the source software
   *
   * @generated from field: optional float visibility = 3;
   */
  visibility?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message TrackedBall.
 * Use `create(TrackedBallSchema)` to create a new message.
 */
export const TrackedBallSchema: GenMessage<TrackedBall, TrackedBallJson> = /*@__PURE__*/
  messageDesc(file_tracked_ssl_vision_detection_tracked, 0);

/**
 * A ball kicked by a robot, including predictions when the ball will come to a stop
 *
 * @generated from message KickedBall
 */
export type KickedBall = Message<"KickedBall"> & {
  /**
   * The initial position [m] from which the ball was kicked
   *
   * @generated from field: required Vector2 pos = 1;
   */
  pos?: Vector2;

  /**
   * The initial velocity [m/s] with which the ball was kicked
   *
   * @generated from field: required Vector3 vel = 2;
   */
  vel?: Vector3;

  /**
   * The unix timestamp [s] when the kick was performed
   *
   * @generated from field: required double start_timestamp = 3;
   */
  startTimestamp: number;

  /**
   * The predicted unix timestamp [s] when the ball comes to a stop
   *
   * @generated from field: optional double stop_timestamp = 4;
   */
  stopTimestamp: number;

  /**
   * The predicted position [m] at which the ball will come to a stop
   *
   * @generated from field: optional Vector2 stop_pos = 5;
   */
  stopPos?: Vector2;

  /**
   * The robot that kicked the ball
   *
   * @generated from field: optional RobotId robot_id = 6;
   */
  robotId?: RobotId;
};

/**
 * A ball kicked by a robot, including predictions when the ball will come to a stop
 *
 * @generated from message KickedBall
 */
export type KickedBallJson = {
  /**
   * The initial position [m] from which the ball was kicked
   *
   * @generated from field: required Vector2 pos = 1;
   */
  pos?: Vector2Json;

  /**
   * The initial velocity [m/s] with which the ball was kicked
   *
   * @generated from field: required Vector3 vel = 2;
   */
  vel?: Vector3Json;

  /**
   * The unix timestamp [s] when the kick was performed
   *
   * @generated from field: required double start_timestamp = 3;
   */
  startTimestamp?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The predicted unix timestamp [s] when the ball comes to a stop
   *
   * @generated from field: optional double stop_timestamp = 4;
   */
  stopTimestamp?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The predicted position [m] at which the ball will come to a stop
   *
   * @generated from field: optional Vector2 stop_pos = 5;
   */
  stopPos?: Vector2Json;

  /**
   * The robot that kicked the ball
   *
   * @generated from field: optional RobotId robot_id = 6;
   */
  robotId?: RobotIdJson;
};

/**
 * Describes the message KickedBall.
 * Use `create(KickedBallSchema)` to create a new message.
 */
export const KickedBallSchema: GenMessage<KickedBall, KickedBallJson> = /*@__PURE__*/
  messageDesc(file_tracked_ssl_vision_detection_tracked, 1);

/**
 * A single tracked robot
 *
 * @generated from message TrackedRobot
 */
export type TrackedRobot = Message<"TrackedRobot"> & {
  /**
   * @generated from field: required RobotId robot_id = 1;
   */
  robotId?: RobotId;

  /**
   * The position [m] in the ssl-vision coordinate system
   *
   * @generated from field: required Vector2 pos = 2;
   */
  pos?: Vector2;

  /**
   * The orientation [rad] in the ssl-vision coordinate system
   *
   * @generated from field: required float orientation = 3;
   */
  orientation: number;

  /**
   * The velocity [m/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional Vector2 vel = 4;
   */
  vel?: Vector2;

  /**
   * The angular velocity [rad/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional float vel_angular = 5;
   */
  velAngular: number;

  /**
   * The visibility of the robot
   * A value between 0 (not visible) and 1 (visible)
   * The exact implementation depends on the source software
   *
   * @generated from field: optional float visibility = 6;
   */
  visibility: number;
};

/**
 * A single tracked robot
 *
 * @generated from message TrackedRobot
 */
export type TrackedRobotJson = {
  /**
   * @generated from field: required RobotId robot_id = 1;
   */
  robotId?: RobotIdJson;

  /**
   * The position [m] in the ssl-vision coordinate system
   *
   * @generated from field: required Vector2 pos = 2;
   */
  pos?: Vector2Json;

  /**
   * The orientation [rad] in the ssl-vision coordinate system
   *
   * @generated from field: required float orientation = 3;
   */
  orientation?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The velocity [m/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional Vector2 vel = 4;
   */
  vel?: Vector2Json;

  /**
   * The angular velocity [rad/s] in the ssl-vision coordinate system
   *
   * @generated from field: optional float vel_angular = 5;
   */
  velAngular?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The visibility of the robot
   * A value between 0 (not visible) and 1 (visible)
   * The exact implementation depends on the source software
   *
   * @generated from field: optional float visibility = 6;
   */
  visibility?: number | "NaN" | "Infinity" | "-Infinity";
};

/**
 * Describes the message TrackedRobot.
 * Use `create(TrackedRobotSchema)` to create a new message.
 */
export const TrackedRobotSchema: GenMessage<TrackedRobot, TrackedRobotJson> = /*@__PURE__*/
  messageDesc(file_tracked_ssl_vision_detection_tracked, 2);

/**
 * A frame that contains all currently tracked objects on the field on all cameras
 *
 * @generated from message TrackedFrame
 */
export type TrackedFrame = Message<"TrackedFrame"> & {
  /**
   * A monotonous increasing frame counter
   *
   * @generated from field: required uint32 frame_number = 1;
   */
  frameNumber: number;

  /**
   * The unix timestamp in [s] of the data
   * If timestamp is larger than timestamp_captured, the source has applied a prediction already
   *
   * @generated from field: required double timestamp = 2;
   */
  timestamp: number;

  /**
   * The list of detected balls
   * The first ball is the primary one
   * Sources may add additional balls based on their capabilities
   *
   * @generated from field: repeated TrackedBall balls = 3;
   */
  balls: TrackedBall[];

  /**
   * The list of detected robots of both teams
   *
   * @generated from field: repeated TrackedRobot robots = 4;
   */
  robots: TrackedRobot[];

  /**
   * Information about a kicked ball, if the ball was kicked by a robot and is still moving
   * Note: This field is optional. Some source implementations might not set this at any time
   *
   * @generated from field: optional KickedBall kicked_ball = 5;
   */
  kickedBall?: KickedBall;

  /**
   * List of capabilities of the source implementation
   *
   * @generated from field: repeated Capability capabilities = 6;
   */
  capabilities: Capability[];
};

/**
 * A frame that contains all currently tracked objects on the field on all cameras
 *
 * @generated from message TrackedFrame
 */
export type TrackedFrameJson = {
  /**
   * A monotonous increasing frame counter
   *
   * @generated from field: required uint32 frame_number = 1;
   */
  frameNumber?: number;

  /**
   * The unix timestamp in [s] of the data
   * If timestamp is larger than timestamp_captured, the source has applied a prediction already
   *
   * @generated from field: required double timestamp = 2;
   */
  timestamp?: number | "NaN" | "Infinity" | "-Infinity";

  /**
   * The list of detected balls
   * The first ball is the primary one
   * Sources may add additional balls based on their capabilities
   *
   * @generated from field: repeated TrackedBall balls = 3;
   */
  balls?: TrackedBallJson[];

  /**
   * The list of detected robots of both teams
   *
   * @generated from field: repeated TrackedRobot robots = 4;
   */
  robots?: TrackedRobotJson[];

  /**
   * Information about a kicked ball, if the ball was kicked by a robot and is still moving
   * Note: This field is optional. Some source implementations might not set this at any time
   *
   * @generated from field: optional KickedBall kicked_ball = 5;
   */
  kickedBall?: KickedBallJson;

  /**
   * List of capabilities of the source implementation
   *
   * @generated from field: repeated Capability capabilities = 6;
   */
  capabilities?: CapabilityJson[];
};

/**
 * Describes the message TrackedFrame.
 * Use `create(TrackedFrameSchema)` to create a new message.
 */
export const TrackedFrameSchema: GenMessage<TrackedFrame, TrackedFrameJson> = /*@__PURE__*/
  messageDesc(file_tracked_ssl_vision_detection_tracked, 3);

/**
 * Capabilities that a source implementation can have
 *
 * @generated from enum Capability
 */
export enum Capability {
  /**
   * @generated from enum value: CAPABILITY_UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * @generated from enum value: CAPABILITY_DETECT_FLYING_BALLS = 1;
   */
  DETECT_FLYING_BALLS = 1,

  /**
   * @generated from enum value: CAPABILITY_DETECT_MULTIPLE_BALLS = 2;
   */
  DETECT_MULTIPLE_BALLS = 2,

  /**
   * @generated from enum value: CAPABILITY_DETECT_KICKED_BALLS = 3;
   */
  DETECT_KICKED_BALLS = 3,
}

/**
 * Capabilities that a source implementation can have
 *
 * @generated from enum Capability
 */
export type CapabilityJson = "CAPABILITY_UNKNOWN" | "CAPABILITY_DETECT_FLYING_BALLS" | "CAPABILITY_DETECT_MULTIPLE_BALLS" | "CAPABILITY_DETECT_KICKED_BALLS";

/**
 * Describes the enum Capability.
 */
export const CapabilitySchema: GenEnum<Capability, CapabilityJson> = /*@__PURE__*/
  enumDesc(file_tracked_ssl_vision_detection_tracked, 0);

