import grpc
import user_pb2
import user_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = user_pb2_grpc.UserServiceStub(channel)

    response = stub.GetUser(user_pb2.UserRequest(id="123"))
    #print(f"User: {response.name}, Email: {response.email}")
    print(response)


if __name__ == '__main__':
    run()
