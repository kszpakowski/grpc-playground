import ProductCatalogue_pb2_grpc
import ProductCatalogue_pb2

class ProductCatalogServicer(ProductCatalogue_pb2_grpc.ProductCatalogServicer):
    def GetProductDetails(self, request, context):
        print(request.id)
        return ProductCatalogue_pb2.ProductDetailsReply(id=request.id, name="test product")
