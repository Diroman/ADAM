import base64
import json
from random import randint, random

import cv2
import grpc
import os
import proto.ml_pb2_grpc
import proto.ml_pb2

from timeit import default_timer as timer


channel = grpc.insecure_channel('34.123.226.60:50051')
stub = proto.ml_pb2_grpc.CarDetectorStub(channel)


def test_predict(path):
    files = os.listdir(path)

    for i in range(len(files)):
        img = cv2.imread(f'files/{files[i]}')
        encoded = base64.b64encode(cv2.imencode('.jpg', img)[1]).decode()
        requestPrediction = proto.ml_pb2.Image(image=encoded)
        responsePrediction = stub.predict(requestPrediction)
        print(files[i], responsePrediction.classes)


if __name__ == '__main__':
    start = timer()
    test_predict('files')
    end = timer()
    print('Done! ', end - start)

