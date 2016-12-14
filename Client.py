#!/usr/bin/python 
import socket 
import json

class PinkTaichi:
    def __init__(self, address):
        self.Address = address 
        self.Port = 2000

    Database = None 
    Collection = None 
    def createConnection(self):
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect((self.Address, self.Port))
        return s
    def sendData(self, data):
        s = self.createConnection()
        s.send(data)
        recv = s.recv(1024)
        return recv
    def insert(self, doc):
        payload = {"Cmd" : "Insert" ,"Database" : self.Database, "Collection" : self.Collection, "Doc" : doc}
        message = json.dumps(payload)
        #s = self.createConnection()
        self.sendData(message)
        #return recv
        #return message
    def findOne(self, doc):
        payload = {"Cmd" : "Find One", "Database" : self.Database, "Collection" : self.Collection, "Doc" : doc}
        message = json.dumps(payload)
        r = self.sendData(message)
        return json.loads(r)
    def find(self, doc): 
        payload = {"Cmd" : "Find", "Database" : self.Database, "Collection" : self.Collection, "Doc" : doc}
        message = json.dumps(payload)
        r = self.sendData(message)
        retVal = json.loads(r)
        return retVal
    def createDatabase(self, database): 
        payload = {"Cmd" : "Create Database", "Database" : database}
        self.sendData(message)
        
    def update(self, doc, newdoc):
        payload = {"Cmd" : "Update", "Database" : self.Database, "Collection" : self.Collection, "Doc" : doc, "NewDoc": newdoc}
        message = json.dumps(payload)
        self.sendData(message)
    def updateOne(self, doc, newdoc): 
        payload = {"Cmd" : "Update One", "Database" : self.Database, "Collection"  : self.Collection, "Doc" : doc, "NewDoc" : newdoc}
        message = json.dumps(payload)
        self.sendData(message)
    def delete(self, doc): 
        payload = {"Cmd" : "Delete", "Database" : self.Database, "Collection" : self.Collection, "Doc" : doc}
        message = json.dumps(payload)
        self.sendData(message)
    def listAll(self):
        payload = {"Cmd" : "List All", "Database" : self.Database, "Collection" : self.Collection}
        payload = json.dumps(payload)
        res = self.sendData(payload)
        retVal = json.loads(res)
        return retVal
    def deleteOne(self, doc): 
        payload = {"Cmd" : "Delete One", "Database" : self.Database, "Collection" : self.Collection, "Doc" : doc}
        message = json.dumps(payload)
        self.sendData(message)
    def showCollections(self):
        payload = {"Cmd" : "Show Collections", "Database" : self.Database}
        message = json.dumps(payload)
        rep = self.sendData(message)
        return json.loads(rep)
    def dropDatabase(self, database):
        payload = {"Cmd" : "Drop Database", "Database" : database}
        message = json.dumps(payload)
        self.sendData(message)
    def createDatabase(self, database):
        payload = {"Cmd" : "Create Database", "Database" : database}
        message = json.dumps(payload)
        self.sendData(message)
    def showDatabases(self):
        payload = {"Cmd" : "Show Databases"}
        message = json.dumps(payload)
        dbList = self.sendData(message)
        retVal = json.loads(dbList)
        return retVal
    def createCollection(self, collection): 
        payload = {"Cmd" : "Create Collection", "Database" : self.Database, "Collection" : collection}
        message = json.dumps(payload)
        self.sendData(message)
    def dropCollection(self, collection):
        payload = {"Cmd" : "Drop Collection", "Database" : self.Database, "Collection" : collection}
        message = json.dumps(payload)
        se = self.sendData(message)

