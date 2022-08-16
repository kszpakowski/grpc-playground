#https://grpc.io/docs/languages/python/basics/


import grpc
from concurrent import futures
import ProductCatalogue_pb2_grpc
import Servicer


def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  ProductCatalogue_pb2_grpc.add_ProductCatalogServicer_to_server(
      Servicer.ProductCatalogServicer(), server)
  server.add_insecure_port('[::]:50051')
  server.start()
  print("Started grpc server")
  server.wait_for_termination()


if __name__ == '__main__':
    serve()