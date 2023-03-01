// source: application_runtime.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.jungletv.ResolveApplicationPageRequest', null, global);
goog.exportSymbol('proto.jungletv.ResolveApplicationPageResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jungletv.ResolveApplicationPageRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.jungletv.ResolveApplicationPageRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.jungletv.ResolveApplicationPageRequest.displayName = 'proto.jungletv.ResolveApplicationPageRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.jungletv.ResolveApplicationPageResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.jungletv.ResolveApplicationPageResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.jungletv.ResolveApplicationPageResponse.displayName = 'proto.jungletv.ResolveApplicationPageResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.jungletv.ResolveApplicationPageRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jungletv.ResolveApplicationPageRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jungletv.ResolveApplicationPageRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    pageId: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jungletv.ResolveApplicationPageRequest}
 */
proto.jungletv.ResolveApplicationPageRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jungletv.ResolveApplicationPageRequest;
  return proto.jungletv.ResolveApplicationPageRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jungletv.ResolveApplicationPageRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jungletv.ResolveApplicationPageRequest}
 */
proto.jungletv.ResolveApplicationPageRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPageId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jungletv.ResolveApplicationPageRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jungletv.ResolveApplicationPageRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jungletv.ResolveApplicationPageRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPageId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string application_id = 1;
 * @return {string}
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.getApplicationId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.jungletv.ResolveApplicationPageRequest} returns this
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.setApplicationId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string page_id = 2;
 * @return {string}
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.getPageId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.jungletv.ResolveApplicationPageRequest} returns this
 */
proto.jungletv.ResolveApplicationPageRequest.prototype.setPageId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.jungletv.ResolveApplicationPageResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.jungletv.ResolveApplicationPageResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jungletv.ResolveApplicationPageResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationFileName: jspb.Message.getFieldWithDefault(msg, 1, ""),
    pageTitle: jspb.Message.getFieldWithDefault(msg, 2, ""),
    applicationVersion: (f = msg.getApplicationVersion()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.jungletv.ResolveApplicationPageResponse}
 */
proto.jungletv.ResolveApplicationPageResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.jungletv.ResolveApplicationPageResponse;
  return proto.jungletv.ResolveApplicationPageResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.jungletv.ResolveApplicationPageResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.jungletv.ResolveApplicationPageResponse}
 */
proto.jungletv.ResolveApplicationPageResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationFileName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPageTitle(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setApplicationVersion(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.jungletv.ResolveApplicationPageResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.jungletv.ResolveApplicationPageResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.jungletv.ResolveApplicationPageResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationFileName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPageTitle();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getApplicationVersion();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string application_file_name = 1;
 * @return {string}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.getApplicationFileName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.jungletv.ResolveApplicationPageResponse} returns this
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.setApplicationFileName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string page_title = 2;
 * @return {string}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.getPageTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.jungletv.ResolveApplicationPageResponse} returns this
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.setPageTitle = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional google.protobuf.Timestamp application_version = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.getApplicationVersion = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.jungletv.ResolveApplicationPageResponse} returns this
*/
proto.jungletv.ResolveApplicationPageResponse.prototype.setApplicationVersion = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.jungletv.ResolveApplicationPageResponse} returns this
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.clearApplicationVersion = function() {
  return this.setApplicationVersion(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.jungletv.ResolveApplicationPageResponse.prototype.hasApplicationVersion = function() {
  return jspb.Message.getField(this, 3) != null;
};


goog.object.extend(exports, proto.jungletv);
