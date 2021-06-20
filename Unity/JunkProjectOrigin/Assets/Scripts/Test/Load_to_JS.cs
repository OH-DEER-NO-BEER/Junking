using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using AOT;
using System;

public class Load_to_JS : MonoBehaviour
{
    [MonoPInvokeCallback(typeof(Action<int>))]
    static void HogeCallback(int fuga)
    {
        Debug.Log(fuga);
    }

    // Start is called before the first frame update
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
