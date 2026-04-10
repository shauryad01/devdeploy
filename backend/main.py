from fastapi import FastAPI

app = FastAPI(title="DevDeploy")

@app.get("/")
def root():
    return {"message": "DevDeploy API running"}