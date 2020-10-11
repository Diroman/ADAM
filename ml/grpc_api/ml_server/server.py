import base64
import grpc
import json
import cv2
import numpy as np
import os
os.chdir("/root/grpc_api/ml_server")
import sys
sys.path.append('/root/grpc_api')
import proto.ml_pb2
import proto.ml_pb2_grpc
from concurrent import futures
from ml_module.car_datector import CarDetector


class PredictServicer(proto.ml_pb2_grpc.CarDetectorServicer):
    def __init__(self):
        with open('config.json') as f:
            config = json.load(f)

        self.model = CarDetector(**config)

    def predict(self, request, context):
        image = request.image
        jpg_original = base64.b64decode(image)
        jpg_as_np = np.frombuffer(jpg_original, dtype=np.uint8)
        image = cv2.imdecode(jpg_as_np, flags=1)
        response = proto.ml_pb2.Classes(classes=self.model.predict(image))
        return response


if __name__ == '__main__':
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    proto.ml_pb2_grpc.add_CarDetectorServicer_to_server(PredictServicer(), server)

    print('Starting server. Listening on port 50051.')
    server.add_insecure_port('0.0.0.0:50051')
    server.start()
    server.wait_for_termination()
