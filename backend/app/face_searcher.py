import pickle
import json
import numpy as np
from scipy.spatial import distance

class face_searcher():
    
    def __init__(self):
        with open('embeddings.pickle', 'rb') as handle:
            self.d = pickle.load(handle)
        self.list_items = []
        for item in list(self.d.values()):
            self.list_items.append(item.squeeze())
        self.keys = json.load( open( "keys_dict.json" ) )
    
    
    def search_face(self, eval_emb, k):
    # Query dataset, k - number of closest elements (returns 2 numpy arrays)
        return_d = {}
        q_labels = distance.cdist(np.expand_dims(eval_emb, axis=0), np.array(self.list_items), 'cosine')
        #print(np.sort(q_labels[0]))
        argmx = np.argsort(q_labels[0])[:k]
        #print(argmx)
        for i in range(len(argmx)):
            return_d['index'] = i
            return_d['dist'] = q_labels[0][argmx[i]]
            return_d['image_url'] = self.keys[str(argmx[i])]   
        return return_d
