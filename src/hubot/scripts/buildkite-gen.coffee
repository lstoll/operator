# Description:
#  Interact with the Buildkite.com Continuous Integration server. Retrieve the  status of projects, setup new ones, and trigger builds.
#
# Commands:
#   hubot buildkite status - List the status of all (i.e. the status of the last build) of one or  all projects.
#   hubot buildkite list-builds - List the last builds of one or all projects, optionally limited to a  branch.
#
# Configuration:
#   OPERATORD_ADDRESS
path = require "path"
grpc = require "grpc"
protobuf = require "protobufjs"

protodir = path.resolve(__dirname + "/../proto")
proto = protobuf.loadProtoFile(root: protodir, file: "buildkite.proto")
buildkite = grpc.loadObject(proto.ns).buildkite
address = process.env.OPERATORD_ADDRESS
if !address
  host = process.env.OPERATORD_PORT_3000_TCP_ADDR
  port = process.env.OPERATORD_PORT_3000_TCP_PORT
  address = "#{host}:#{port}"

client = new buildkite.BuildkiteService(address, grpc.Credentials.createInsecure())

module.exports = (robot) ->

  robot.respond /buildkite status slug=(\w+)/, (msg) ->
    client.status {slug: msg.match[1],}, (err, response) ->
      if err
        msg.send("```\nStatus error: #{err.message}\n```")
      else
        msg.send("```\n#{response.output.PlainText}\n```")

  robot.respond /buildkite list-builds project_slug=(\w+)/, (msg) ->
    client.listBuilds {project_slug: msg.match[1],}, (err, response) ->
      if err
        msg.send("```\nListBuilds error: #{err.message}\n```")
      else
        msg.send("```\n#{response.output.PlainText}\n```")

