from flask import Flask, escape, request

import grpc

import shortn_pb2
import shortn_pb2_grpc

ADDR = "localhost:52002"
app = Flask(__name__)


@app.route('/')
def index():
    return "Welcome to shortn url."


@app.route('/api/add_page', methods=['GET'])
def add_page():
    raw_url = escape(request.args.get('url', ''))
    if len(raw_url) == 0:
        return "/api/add_page?url={{ DST URL }}"

    try:
        with grpc.insecure_channel(ADDR) as channel:
            stub = shortn_pb2_grpc.ShortnStub(channel)
            response = stub.AddPageRpc(
                shortn_pb2.AddPageRequest(raw_url=raw_url))
    except grpc.RpcError as e:
        return f"RPC error: {e}", 500
    except Exception as e:
        return f"Unexpected error: {e}", 500
    return f"{response.short_url} -> {raw_url}"
