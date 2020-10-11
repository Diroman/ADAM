import gc
import torch
import numpy as np
import torch.nn as nn
import torch.nn.functional as F
import torchvision.models as models
import torchvision.transforms as transforms
# import proto.ml_pb2_grpc
# import proto.ml_pb2
import PIL.Image as Image


class CarDetector:
    def __init__(self, input_shape, path_to_model_cpu, path_to_model_cuda):
        self.device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")
        self.shape = input_shape
        self.model = self.build_model(path_to_model_cpu, path_to_model_cuda, self.device)
        self.model.eval()
        self.loader = transforms.Compose([transforms.Resize((self.shape, self.shape)),
                                          transforms.ToTensor(),
                                          transforms.Normalize((0.5, 0.5, 0.5), (0.5, 0.5, 0.5))])
        self.classes = ['Hyundai Solaris sedan',
                        'KIA Rio sedan',
                        'SKODA OCTAVIA sedan',
                        'Volkswagen Polo sedan',
                        'Volkswagen Tiguan']

    @staticmethod
    def build_model(path_to_model_cpu, path_to_model_cuda, device):
        model = models.resnet34()
        num_filters = model.fc.in_features
        model.fc = nn.Linear(num_filters, 5)
        checkpoint = torch.load(path_to_model_cpu if device.type == 'cpu' else path_to_model_cuda)
        model.load_state_dict(checkpoint['state_dict'])
        model = model.to(device)
        return model

    def predict(self, image):
        image = Image.fromarray(np.uint8(image)).convert('RGB')
        image = self.loader(image).float()
        image = torch.autograd.Variable(image, requires_grad=True)
        image = image.unsqueeze(0)
        image = image.to(self.device)
        with torch.no_grad():
            data = self.model(image)
            output = F.softmax(data)
        result = {self.classes[i]: value.item() for i, value in enumerate(output.data[0])}
        del output
        del data
        del image
        gc.collect()
        return result
