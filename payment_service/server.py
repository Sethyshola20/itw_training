import os
import time
import grpc
from concurrent import futures

from grpc_tools import protoc  # type: ignore

proto_path = os.path.join(os.path.dirname(__file__), 'payment.proto')
out_dir = os.path.dirname(__file__)
protoc.main([
    '',
    f'-I{os.path.dirname(proto_path)}',
    f'--python_out={out_dir}',
    f'--grpc_python_out={out_dir}',
    proto_path,
])

import payment_pb2  # type: ignore
import payment_pb2_grpc  # type: ignore

class PaymentService(payment_pb2_grpc.PaymentServiceServicer):
    def RecordPayment(self, request, context):
        return payment_pb2.PaymentResponse(payment_id="pay_" + str(int(time.time())), status="success")

    def ValidatePayment(self, request, context):
        return payment_pb2.ValidateResponse(payment_id=request.payment_id, status="success")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    payment_pb2_grpc.add_PaymentServiceServicer_to_server(PaymentService(), server)
    port = os.getenv('PAYMENT_GRPC_PORT', '50051')
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()