# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import ProductCatalogue_pb2 as ProductCatalogue__pb2


class ProductCatalogStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetProductDetails = channel.unary_unary(
                '/product.ProductCatalog/GetProductDetails',
                request_serializer=ProductCatalogue__pb2.ProductDetailsRequest.SerializeToString,
                response_deserializer=ProductCatalogue__pb2.ProductDetailsReply.FromString,
                )


class ProductCatalogServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetProductDetails(self, request, context):
        """Get product details
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProductCatalogServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetProductDetails': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProductDetails,
                    request_deserializer=ProductCatalogue__pb2.ProductDetailsRequest.FromString,
                    response_serializer=ProductCatalogue__pb2.ProductDetailsReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'product.ProductCatalog', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProductCatalog(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetProductDetails(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/product.ProductCatalog/GetProductDetails',
            ProductCatalogue__pb2.ProductDetailsRequest.SerializeToString,
            ProductCatalogue__pb2.ProductDetailsReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
