public class TestClass {
    TestClass() {
        vboolean = true;
        vbyte = 10;
        vchar = 10;
        vshort = 10;
        vint = 10;
        vlong = 10;
        vfloat = 10;
        vdouble = 10;
        vclass = new TestSubClass();

        aboolean = new boolean[]{true, false};
        abyte = new byte[]{10, 11};
        achar = new char[]{10, 11};
        ashort = new short[]{10, 11};
        aint = new int[]{10, 11};
        along = new long[]{10, 11};
        afloat = new float[]{10, 11};
        adouble = new double[]{10, 11};
        aclass = new TestSubClass[]{vclass, vclass};
    }

    TestClass(int a) {
    }

    public boolean vboolean;
    public byte vbyte;
    public char vchar;
    public short vshort;
    public int vint;
    public long vlong;
    public float vfloat;
    public double vdouble;
    public TestSubClass vclass;

    public boolean[] aboolean;
    public byte[] abyte;
    public char[] achar;
    public short[] ashort;
    public int[] aint;
    public long[] along;
    public float[] afloat;
    public double[] adouble;
    public TestSubClass[] aclass;

    public static boolean svboolean = false;
    public static byte svbyte = 10;
    public static char svchar = 10;
    public static short svshort = 10;
    public static int svint = 10;
    public static long svlong = 10;
    public static float svfloat = 10;
    public static double svdouble = 10;
    public static TestSubClass svclass;

    public static boolean[] saboolean = new boolean[]{true, false};
    public static byte[] sabyte = new byte[]{10, 11};
    public static char[] sachar = new char[]{10, 11};
    public static short[] sashort = new short[]{10, 11};
    public static int[] saint = new int[]{10, 11};
    public static long[] salong = new long[]{10, 11};
    public static float[] safloat = new float[]{10, 11};
    public static double[] sadouble = new double[]{10, 11};
    public static TestSubClass[] saclass = new TestSubClass[]{};

    public boolean mvboolean(){ return vboolean; }
    public byte mvbyte(){ return vbyte; }
    public char mvchar(){ return vchar; }
    public short mvshort(){ return vshort; }
    public int mvint(){ return vint; }
    public long mvlong(){ return vlong; }
    public float mvfloat(){ return vfloat; }
    public double mvdouble(){ return vdouble; }
    public TestSubClass mvclass(){ return vclass; }

    public boolean[] maboolean() { return aboolean; }
    public byte[] mabyte() { return abyte; }
    public char[] machar() { return achar; }
    public short[] mashort() { return ashort; }
    public int[] maint() { return aint; }
    public long[] malong() { return along; }
    public float[] mafloat() { return afloat; }
    public double[] madouble() { return adouble; }
    public TestSubClass[] maclass() { return aclass; }

    public static boolean smvboolean(){ return svboolean; }
    public static byte smvbyte(){ return svbyte; }
    public static char smvchar(){ return svchar; }
    public static short smvshort(){ return svshort; }
    public static int smvint(){ return svint; }
    public static long smvlong(){ return svlong; }
    public static float smvfloat(){ return svfloat; }
    public static double smvdouble(){ return svdouble; }
    public static TestSubClass smvclass(){ return svclass; }

    public static boolean[] smaboolean() { return saboolean; }
    public static byte[] smabyte() { return sabyte; }
    public static char[] smachar() { return sachar; }
    public static short[] smashort() { return sashort; }
    public static int[] smaint() { return saint; }
    public static long[] smalong() { return salong; }
    public static float[] smafloat() { return safloat; }
    public static double[] smadouble() { return sadouble; }
    public static TestSubClass[] smaclass() { return saclass; }
}

