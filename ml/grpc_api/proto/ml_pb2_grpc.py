# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import proto.ml_pb2 as ml__pb2


class CarDetectorStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.predict = channel.unary_unary(
                '/CarDetector/predict',
                request_serializer=ml__pb2.Image.SerializeToString,
                response_deserializer=ml__pb2.Classes.FromString,
                )


class CarDetectorServicer(object):
    """Missing associated documentation comment in .proto file"""

    def predict(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CarDetectorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'predict': grpc.unary_unary_rpc_method_handler(
                    servicer.predict,
                    request_deserializer=ml__pb2.Image.FromString,
                    response_serializer=ml__pb2.Classes.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'CarDetector', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CarDetector(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def predict(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/CarDetector/predict',
            ml__pb2.Image.SerializeToString,
            ml__pb2.Classes.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
